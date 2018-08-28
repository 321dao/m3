// Copyright (c) 2018 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package logical

import (
	"errors"

	"github.com/m3db/m3/src/query/block"
	"github.com/m3db/m3/src/query/models"
)

// VectorMatchCardinality describes the cardinality relationship
// of two Vectors in a binary operation.
type VectorMatchCardinality int

const (
	// CardOneToOne is used for one-one relationship
	CardOneToOne VectorMatchCardinality = iota
	// CardManyToOne is used for many-one relationship
	CardManyToOne
	// CardOneToMany is used for one-many relationship
	CardOneToMany
	// CardManyToMany is used for many-many relationship
	CardManyToMany
)

// VectorMatching describes how elements from two Vectors in a binary
// operation are supposed to be matched.
type VectorMatching struct {
	// The cardinality of the two Vectors.
	Card VectorMatchCardinality
	// MatchingLabels contains the labels which define equality of a pair of
	// elements from the Vectors.
	MatchingLabels []string
	// On includes the given label names from matching,
	// rather than excluding them.
	On bool
	// Include contains additional labels that should be included in
	// the result from the side with the lower cardinality.
	Include []string
}

// HashFunc returns a function that calculates the signature for a metric
// ignoring the provided labels. If on, then the given labels are only used instead.
func HashFunc(on bool, names ...string) func(models.Tags) uint64 {
	if on {
		return func(tags models.Tags) uint64 { return tags.IDWithKeys(names...) }
	}
	return func(tags models.Tags) uint64 { return tags.IDWithExcludes(names...) }
}

const initIndexSliceLength = 10

var (
	errMismatchedBounds     = errors.New("block bounds are mismatched")
	errMismatchedStepCounts = errors.New("block step counts are mismatched")
)

func combineMetaAndSeriesMeta(
	meta, otherMeta block.Metadata,
	seriesMeta, otherSeriesMeta []block.SeriesMeta,
) (block.Metadata, []block.SeriesMeta, []block.SeriesMeta, error) {
	if !meta.Bounds.Equals(otherMeta.Bounds) {
		return block.Metadata{},
			[]block.SeriesMeta{},
			[]block.SeriesMeta{},
			errMismatchedBounds
	}

	// NB (arnikola): mutating tags in `meta` to avoid allocations
	leftTags := meta.Tags
	otherTags := otherMeta.Tags.TagMap()

	metaTagsToAdd := make(models.Tags, 0)
	otherMetaTagsToAdd := make(models.Tags, 0)
	tags := make(models.Tags, 0)

	for _, t := range leftTags {
		if otherTag, ok := otherTags[t.Name]; ok {
			if t.Value != otherTag.Value {
				// If both metas have the same common tag  with different
				// values, remove it from common tag list and explicitly
				// add it to each seriesMeta.
				metaTagsToAdd = append(metaTagsToAdd, t)
				otherMetaTagsToAdd = append(otherMetaTagsToAdd, otherTag)
			} else {
				tags = append(tags, t)
			}

			// NB(arnikola): delete common tag from otherTags as it
			// has already been handled
			delete(otherTags, t.Name)
		} else {
			// Tag does not exist on otherMeta; remove it
			// from common tags and explicitly add it to each seriesMeta
			metaTagsToAdd = append(metaTagsToAdd, t)
		}
	}

	// Iterate over otherMeta common tags and explicitly add
	// remaining tags to otherSeriesMeta
	for _, otherTag := range otherTags {
		otherMetaTagsToAdd = append(otherMetaTagsToAdd, otherTag)
	}

	// Set common tags
	meta.Tags = tags
	for i, meta := range seriesMeta {
		seriesMeta[i].Tags = meta.Tags.Add(metaTagsToAdd)
	}

	for i, meta := range otherSeriesMeta {
		otherSeriesMeta[i].Tags = meta.Tags.Add(otherMetaTagsToAdd)
	}

	return meta,
		seriesMeta,
		otherSeriesMeta,
		nil
}

// FlattenMetadata applies all shared tags from Metadata to each SeriesMeta
func FlattenMetadata(
	meta block.Metadata,
	seriesMeta []block.SeriesMeta,
) []block.SeriesMeta {
	for k, v := range meta.Tags {
		for _, metas := range seriesMeta {
			metas.Tags[k] = v
		}
	}

	return seriesMeta
}

// DedupeMetadata applies all shared tags from Metadata to each SeriesMeta
func DedupeMetadata(
	seriesMeta []block.SeriesMeta,
) (models.Tags, []block.SeriesMeta) {
	if len(seriesMeta) == 0 {
		return make(models.Tags, 0), seriesMeta
	}

	commonKeys := make([]string, 0, len(seriesMeta[0].Tags))
	commonTags := make(map[string]string, len(seriesMeta[0].Tags))
	// For each tag in the first series, read through list of seriesMetas;
	// if key not found or value differs, this is not a shared tag
	var distinct bool
	for _, t := range seriesMeta[0].Tags {
		distinct = false
		for _, metas := range seriesMeta[1:] {
			if val, ok := metas.Tags.Get(t.Name); ok {
				if val != t.Value {
					distinct = true
					break
				}
			} else {
				distinct = true
				break
			}
		}

		if !distinct {
			// This is a shared tag; add it to shared meta
			commonKeys = append(commonKeys, t.Name)
			commonTags[t.Name] = t.Value
		}
	}

	for i, meta := range seriesMeta {
		seriesMeta[i].Tags = meta.Tags.WithoutKeys(commonKeys)
	}

	return models.FromMap(commonTags), seriesMeta
}

func appendValuesAtIndices(idxArray []int, iter block.StepIter, builder block.Builder) error {
	for index := 0; iter.Next(); index++ {
		step, err := iter.Current()
		if err != nil {
			return err
		}

		values := step.Values()
		for _, idx := range idxArray {
			builder.AppendValue(index, values[idx])
		}
	}

	return nil
}
