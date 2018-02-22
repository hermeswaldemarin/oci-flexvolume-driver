// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// File Storage Service API
//
// APIs for OCI file storage service.
//

package ffsw

import (
    "github.com/oracle/oci-go-sdk/common"
)


    
 // Snapshot A point in time snapshot of a specified file system.
type Snapshot struct {
    
 // The OCID of the snapshot's file system.
    FileSystemId *string `mandatory:"true" json:"fileSystemId"`
    
 // The OCID of the snapshot.
    Id *string `mandatory:"true" json:"id"`
    
 // The current state of the snapshot.
    LifecycleState SnapshotLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`
    
 // Name of the snapshot. Immutable.
 // Avoid entering confidential information.
 // Example: `sunday`
    Name *string `mandatory:"true" json:"name"`
    
 // The date and time the snapshot was created, in the format defined by RFC3339.
 // Example: `2016-08-25T21:10:29.600Z`
    TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`
}

func (m Snapshot) String() string {
    return common.PointerString(m)
}


// SnapshotLifecycleStateEnum Enum with underlying type: string
type SnapshotLifecycleStateEnum string

// Set of constants representing the allowable values for SnapshotLifecycleState
const (
    SnapshotLifecycleStateCreating SnapshotLifecycleStateEnum = "CREATING"
    SnapshotLifecycleStateActive SnapshotLifecycleStateEnum = "ACTIVE"
    SnapshotLifecycleStateDeleting SnapshotLifecycleStateEnum = "DELETING"
    SnapshotLifecycleStateDeleted SnapshotLifecycleStateEnum = "DELETED"
    SnapshotLifecycleStateUnknown SnapshotLifecycleStateEnum = "UNKNOWN"
)

var mappingSnapshotLifecycleState = map[string]SnapshotLifecycleStateEnum { 
    "CREATING": SnapshotLifecycleStateCreating,
    "ACTIVE": SnapshotLifecycleStateActive,
    "DELETING": SnapshotLifecycleStateDeleting,
    "DELETED": SnapshotLifecycleStateDeleted,
    "UNKNOWN": SnapshotLifecycleStateUnknown,
}

// GetSnapshotLifecycleStateEnumValues Enumerates the set of values for SnapshotLifecycleState
func GetSnapshotLifecycleStateEnumValues() []SnapshotLifecycleStateEnum {
   values := make([]SnapshotLifecycleStateEnum, 0)
   for _, v := range mappingSnapshotLifecycleState {
      if v != SnapshotLifecycleStateUnknown {
         values = append(values, v)
      }
   }
   return values
}


