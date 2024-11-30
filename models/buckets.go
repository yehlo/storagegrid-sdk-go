package models

import "time"

type Bucket struct {
	// if true, object versioning will be enabled for the bucket.
	EnableVersioning *bool `json:"enableVersioning,omitempty"`
	// Bucket name. Must be unique across the grid and DNS compliant. See the instructions for using S3 for details.
	Name string `json:"name"`
	// the region for this bucket, which must already be defined (defaults to us-east-1 if not specified)
	Region       string                      `json:"region,omitempty"`
	S3ObjectLock *BucketS3ObjectLockSettings `json:"s3ObjectLock,omitempty"`
	// the creation time of the bucket
	CreationTime time.Time `json:"creationTime,omitempty"`
	// compliance settings for the bucket
	Compliance *BucketComplianceSettings `json:"compliance,omitempty"`
	// status of the object deletion
	DeleteObjectStatus *BucketDeleteObjectStatus `json:"deleteObjectStatus,omitempty"`
}

// BucketS3ObjectLockSettings Settings for S3 Object Lock. Cannot be used with legacy Compliance.
type BucketS3ObjectLockSettings struct {
	// whether the bucket has S3 Object Lock enabled
	Enabled                 *bool                                       `json:"enabled"`
	DefaultRetentionSetting *BucketS3ObjectLockDefaultRetentionSettings `json:"defaultRetentionSetting,omitempty"`
}

// BucketS3ObjectLockDefaultRetentionSettings Default retention settings for S3 Object Lock.
type BucketS3ObjectLockDefaultRetentionSettings struct {
	// The retention mode used for new objects added to this bucket. Must be compliance, which means that an object version cannot be overwritten or deleted by any user.
	Mode string `json:"mode"`
	// The length of the default retention period for new objects added to this bucket, in days. If provided, must be paired with retentionMode. Does not affect existing bucket objects or objects with their own retain-until-date settings.
	Days int32 `json:"days,omitempty"`
	// The length of the default retention period for new objects added to this bucket, in years. If provided, must be paired with retentionMode. Does not affect existing bucket objects or objects with their own retain-until-date settings.
	Years int32 `json:"years,omitempty"`
}

type BucketComplianceSettings struct {
	// Wether the objects are autoDeleted
	AutoDelete *bool `json:"autoDelete"`
	// time to legally hold the objects
	LegalHold *bool `json:"legalHold"`
	// amount of minuts for retention
	RetentionPeriodMinutes *int32 `json:"retentionPeriodMinutes,omitempty"`
}

type BucketDeleteObjectStatus struct {
	// are the objects being deleted
	IsDeletingObjects *bool `json:"isDeletingObjects"`
	// initial Object count before operation
	InitialObjectCount *int32 `json:"initialObjectCount,omitempty"`
	// initial Object Bytes before the operation
	InitialObjectBytes *int64 `json:"initialObjectBytes,omitempty"`
}
