// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.


// Code generated from the elasticsearch-specification DO NOT EDIT.
// https://github.com/elastic/elasticsearch-specification/tree/135ae054e304239743b5777ad8d41cb2c9091d35


package types

// BucketsIpRangeBucket holds the union for the following types:
//     []IpRangeBucket
//     map[string]IpRangeBucket
//
// https://github.com/elastic/elasticsearch-specification/blob/135ae054e304239743b5777ad8d41cb2c9091d35/specification/_types/aggregations/Aggregate.ts#L303-L312
type BucketsIpRangeBucket interface{}

// BucketsIpRangeBucketBuilder holds BucketsIpRangeBucket struct and provides a builder API.
type BucketsIpRangeBucketBuilder struct {
	v BucketsIpRangeBucket
}

// NewBucketsIpRangeBucket provides a builder for the BucketsIpRangeBucket struct.
func NewBucketsIpRangeBucketBuilder() *BucketsIpRangeBucketBuilder {
	return &BucketsIpRangeBucketBuilder{}
}

// Build finalize the chain and returns the BucketsIpRangeBucket struct
func (u *BucketsIpRangeBucketBuilder) Build() BucketsIpRangeBucket {
	return u.v
}

func (u *BucketsIpRangeBucketBuilder) IpRangeBuckets(iprangebuckets []IpRangeBucketBuilder) *BucketsIpRangeBucketBuilder {
	tmp := make([]IpRangeBucket, len(iprangebuckets))
	for _, value := range iprangebuckets {
		tmp = append(tmp, value.Build())
	}
	u.v = tmp
	return u
}

func (u *BucketsIpRangeBucketBuilder) Map(values map[string]*IpRangeBucketBuilder) *BucketsIpRangeBucketBuilder {
	tmp := make(map[string]IpRangeBucket, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	u.v = tmp
	return u
}