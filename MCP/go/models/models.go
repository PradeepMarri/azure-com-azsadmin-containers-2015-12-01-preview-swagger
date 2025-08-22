package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// MigrationParameters represents the MigrationParameters schema from the OpenAPI specification
type MigrationParameters struct {
	Destinationshareuncpath string `json:"destinationShareUncPath"` // The UNC path of the destination share for migration.
	Storageaccountname string `json:"storageAccountName"` // The name of the storage account where the container locates.
	Containername string `json:"containerName"` // The name of the container to be migrated.
}

// MigrationResult represents the MigrationResult schema from the OpenAPI specification
type MigrationResult struct {
	Sourcesharename string `json:"sourceShareName,omitempty"` // The name of the source storage share.
	Destinationsharename string `json:"destinationShareName,omitempty"` // The name of the destination storage share.
	Containername string `json:"containerName,omitempty"` // The name of the container to be migrated.
	Migrationstatus string `json:"migrationStatus,omitempty"`
	Storageaccountname string `json:"storageAccountName,omitempty"` // The storage account name.
	Subentitiescompleted int64 `json:"subEntitiesCompleted,omitempty"` // The number of entities which have been migrated.
	Failurereason string `json:"failureReason,omitempty"` // The migration failure reason.
	Subentitiesfailed int64 `json:"subEntitiesFailed,omitempty"` // The number of entities which failed in migration.
	Jobid string `json:"jobId,omitempty"` // The migration job ID.
}

// Container represents the Container schema from the OpenAPI specification
type Container struct {
	Containerstate string `json:"containerState,omitempty"` // The current state of the container.
	Containerid int64 `json:"containerid,omitempty"` // The container ID.
	Containername string `json:"containername,omitempty"` // Container name.
	Sharename string `json:"sharename,omitempty"` // The name of the share where the container locates.
	Usedbytesinprimaryvolume int64 `json:"usedBytesInPrimaryVolume,omitempty"` // The used space, in bytes, of the container in the primary volume.
	Accountid string `json:"accountid,omitempty"` // The ID of the storage account.
	Accountname string `json:"accountname,omitempty"` // The name of storage account where the container locates.
}
