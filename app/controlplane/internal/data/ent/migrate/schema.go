// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// APITokensColumns holds the columns for the "api_tokens" table.
	APITokensColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "name", Type: field.TypeString},
		{Name: "description", Type: field.TypeString, Nullable: true},
		{Name: "created_at", Type: field.TypeTime, Default: "CURRENT_TIMESTAMP"},
		{Name: "expires_at", Type: field.TypeTime, Nullable: true},
		{Name: "revoked_at", Type: field.TypeTime, Nullable: true},
		{Name: "organization_id", Type: field.TypeUUID},
	}
	// APITokensTable holds the schema information for the "api_tokens" table.
	APITokensTable = &schema.Table{
		Name:       "api_tokens",
		Columns:    APITokensColumns,
		PrimaryKey: []*schema.Column{APITokensColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "api_tokens_organizations_api_tokens",
				Columns:    []*schema.Column{APITokensColumns[6]},
				RefColumns: []*schema.Column{OrganizationsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "apitoken_name_organization_id",
				Unique:  true,
				Columns: []*schema.Column{APITokensColumns[1], APITokensColumns[6]},
				Annotation: &entsql.IndexAnnotation{
					Where: "revoked_at IS NULL",
				},
			},
		},
	}
	// CasBackendsColumns holds the columns for the "cas_backends" table.
	CasBackendsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "location", Type: field.TypeString},
		{Name: "name", Type: field.TypeString},
		{Name: "provider", Type: field.TypeEnum, Enums: []string{"AzureBlob", "OCI", "INLINE", "AWS-S3"}},
		{Name: "description", Type: field.TypeString, Nullable: true},
		{Name: "secret_name", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime, Default: "CURRENT_TIMESTAMP"},
		{Name: "validation_status", Type: field.TypeEnum, Enums: []string{"OK", "Invalid"}, Default: "OK"},
		{Name: "validated_at", Type: field.TypeTime, Default: "CURRENT_TIMESTAMP"},
		{Name: "default", Type: field.TypeBool, Default: false},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "fallback", Type: field.TypeBool, Default: false},
		{Name: "organization_cas_backends", Type: field.TypeUUID},
	}
	// CasBackendsTable holds the schema information for the "cas_backends" table.
	CasBackendsTable = &schema.Table{
		Name:       "cas_backends",
		Columns:    CasBackendsColumns,
		PrimaryKey: []*schema.Column{CasBackendsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "cas_backends_organizations_cas_backends",
				Columns:    []*schema.Column{CasBackendsColumns[12]},
				RefColumns: []*schema.Column{OrganizationsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "casbackend_name_organization_cas_backends",
				Unique:  true,
				Columns: []*schema.Column{CasBackendsColumns[2], CasBackendsColumns[12]},
				Annotation: &entsql.IndexAnnotation{
					Where: "deleted_at IS NULL",
				},
			},
		},
	}
	// CasMappingsColumns holds the columns for the "cas_mappings" table.
	CasMappingsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "digest", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime, Default: "CURRENT_TIMESTAMP"},
		{Name: "cas_mapping_cas_backend", Type: field.TypeUUID},
		{Name: "cas_mapping_workflow_run", Type: field.TypeUUID, Nullable: true},
		{Name: "cas_mapping_organization", Type: field.TypeUUID},
	}
	// CasMappingsTable holds the schema information for the "cas_mappings" table.
	CasMappingsTable = &schema.Table{
		Name:       "cas_mappings",
		Columns:    CasMappingsColumns,
		PrimaryKey: []*schema.Column{CasMappingsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "cas_mappings_cas_backends_cas_backend",
				Columns:    []*schema.Column{CasMappingsColumns[3]},
				RefColumns: []*schema.Column{CasBackendsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "cas_mappings_workflow_runs_workflow_run",
				Columns:    []*schema.Column{CasMappingsColumns[4]},
				RefColumns: []*schema.Column{WorkflowRunsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "cas_mappings_organizations_organization",
				Columns:    []*schema.Column{CasMappingsColumns[5]},
				RefColumns: []*schema.Column{OrganizationsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "casmapping_digest",
				Unique:  false,
				Columns: []*schema.Column{CasMappingsColumns[1]},
			},
		},
	}
	// IntegrationsColumns holds the columns for the "integrations" table.
	IntegrationsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "name", Type: field.TypeString},
		{Name: "kind", Type: field.TypeString},
		{Name: "description", Type: field.TypeString, Nullable: true},
		{Name: "secret_name", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime, Default: "CURRENT_TIMESTAMP"},
		{Name: "configuration", Type: field.TypeBytes, Nullable: true},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "organization_integrations", Type: field.TypeUUID},
	}
	// IntegrationsTable holds the schema information for the "integrations" table.
	IntegrationsTable = &schema.Table{
		Name:       "integrations",
		Columns:    IntegrationsColumns,
		PrimaryKey: []*schema.Column{IntegrationsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "integrations_organizations_integrations",
				Columns:    []*schema.Column{IntegrationsColumns[8]},
				RefColumns: []*schema.Column{OrganizationsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "integration_name_organization_integrations",
				Unique:  true,
				Columns: []*schema.Column{IntegrationsColumns[1], IntegrationsColumns[8]},
				Annotation: &entsql.IndexAnnotation{
					Where: "deleted_at IS NULL",
				},
			},
		},
	}
	// IntegrationAttachmentsColumns holds the columns for the "integration_attachments" table.
	IntegrationAttachmentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeTime, Default: "CURRENT_TIMESTAMP"},
		{Name: "configuration", Type: field.TypeBytes, Nullable: true},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "integration_attachment_integration", Type: field.TypeUUID},
		{Name: "integration_attachment_workflow", Type: field.TypeUUID},
	}
	// IntegrationAttachmentsTable holds the schema information for the "integration_attachments" table.
	IntegrationAttachmentsTable = &schema.Table{
		Name:       "integration_attachments",
		Columns:    IntegrationAttachmentsColumns,
		PrimaryKey: []*schema.Column{IntegrationAttachmentsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "integration_attachments_integrations_integration",
				Columns:    []*schema.Column{IntegrationAttachmentsColumns[4]},
				RefColumns: []*schema.Column{IntegrationsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "integration_attachments_workflows_workflow",
				Columns:    []*schema.Column{IntegrationAttachmentsColumns[5]},
				RefColumns: []*schema.Column{WorkflowsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// MembershipsColumns holds the columns for the "memberships" table.
	MembershipsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "current", Type: field.TypeBool, Default: false},
		{Name: "created_at", Type: field.TypeTime, Default: "CURRENT_TIMESTAMP"},
		{Name: "updated_at", Type: field.TypeTime, Default: "CURRENT_TIMESTAMP"},
		{Name: "role", Type: field.TypeEnum, Enums: []string{"role:org:owner", "role:org:admin", "role:org:viewer"}},
		{Name: "organization_memberships", Type: field.TypeUUID},
		{Name: "user_memberships", Type: field.TypeUUID},
	}
	// MembershipsTable holds the schema information for the "memberships" table.
	MembershipsTable = &schema.Table{
		Name:       "memberships",
		Columns:    MembershipsColumns,
		PrimaryKey: []*schema.Column{MembershipsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "memberships_organizations_memberships",
				Columns:    []*schema.Column{MembershipsColumns[5]},
				RefColumns: []*schema.Column{OrganizationsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "memberships_users_memberships",
				Columns:    []*schema.Column{MembershipsColumns[6]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "membership_organization_memberships_user_memberships",
				Unique:  true,
				Columns: []*schema.Column{MembershipsColumns[5], MembershipsColumns[6]},
			},
		},
	}
	// OrgInvitationsColumns holds the columns for the "org_invitations" table.
	OrgInvitationsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "receiver_email", Type: field.TypeString},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"accepted", "pending"}, Default: "pending"},
		{Name: "created_at", Type: field.TypeTime, Default: "CURRENT_TIMESTAMP"},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "role", Type: field.TypeEnum, Nullable: true, Enums: []string{"role:org:owner", "role:org:admin", "role:org:viewer"}},
		{Name: "organization_id", Type: field.TypeUUID},
		{Name: "sender_id", Type: field.TypeUUID},
	}
	// OrgInvitationsTable holds the schema information for the "org_invitations" table.
	OrgInvitationsTable = &schema.Table{
		Name:       "org_invitations",
		Columns:    OrgInvitationsColumns,
		PrimaryKey: []*schema.Column{OrgInvitationsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "org_invitations_organizations_organization",
				Columns:    []*schema.Column{OrgInvitationsColumns[6]},
				RefColumns: []*schema.Column{OrganizationsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "org_invitations_users_sender",
				Columns:    []*schema.Column{OrgInvitationsColumns[7]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// OrganizationsColumns holds the columns for the "organizations" table.
	OrganizationsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "created_at", Type: field.TypeTime, Default: "CURRENT_TIMESTAMP"},
	}
	// OrganizationsTable holds the schema information for the "organizations" table.
	OrganizationsTable = &schema.Table{
		Name:       "organizations",
		Columns:    OrganizationsColumns,
		PrimaryKey: []*schema.Column{OrganizationsColumns[0]},
	}
	// ReferrersColumns holds the columns for the "referrers" table.
	ReferrersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "digest", Type: field.TypeString},
		{Name: "kind", Type: field.TypeString},
		{Name: "downloadable", Type: field.TypeBool},
		{Name: "created_at", Type: field.TypeTime, Default: "CURRENT_TIMESTAMP"},
		{Name: "metadata", Type: field.TypeJSON, Nullable: true},
		{Name: "annotations", Type: field.TypeJSON, Nullable: true},
	}
	// ReferrersTable holds the schema information for the "referrers" table.
	ReferrersTable = &schema.Table{
		Name:       "referrers",
		Columns:    ReferrersColumns,
		PrimaryKey: []*schema.Column{ReferrersColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "referrer_digest_kind",
				Unique:  true,
				Columns: []*schema.Column{ReferrersColumns[1], ReferrersColumns[2]},
			},
		},
	}
	// RobotAccountsColumns holds the columns for the "robot_accounts" table.
	RobotAccountsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "name", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime, Default: "CURRENT_TIMESTAMP"},
		{Name: "revoked_at", Type: field.TypeTime, Nullable: true},
		{Name: "workflow_robotaccounts", Type: field.TypeUUID, Nullable: true},
	}
	// RobotAccountsTable holds the schema information for the "robot_accounts" table.
	RobotAccountsTable = &schema.Table{
		Name:       "robot_accounts",
		Columns:    RobotAccountsColumns,
		PrimaryKey: []*schema.Column{RobotAccountsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "robot_accounts_workflows_robotaccounts",
				Columns:    []*schema.Column{RobotAccountsColumns[4]},
				RefColumns: []*schema.Column{WorkflowsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "created_at", Type: field.TypeTime, Default: "CURRENT_TIMESTAMP"},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// WorkflowsColumns holds the columns for the "workflows" table.
	WorkflowsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "name", Type: field.TypeString},
		{Name: "project", Type: field.TypeString, Nullable: true},
		{Name: "team", Type: field.TypeString, Nullable: true},
		{Name: "runs_count", Type: field.TypeInt, Default: 0},
		{Name: "created_at", Type: field.TypeTime, Default: "CURRENT_TIMESTAMP"},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "public", Type: field.TypeBool, Default: false},
		{Name: "description", Type: field.TypeString, Nullable: true},
		{Name: "organization_id", Type: field.TypeUUID},
		{Name: "workflow_contract", Type: field.TypeUUID},
	}
	// WorkflowsTable holds the schema information for the "workflows" table.
	WorkflowsTable = &schema.Table{
		Name:       "workflows",
		Columns:    WorkflowsColumns,
		PrimaryKey: []*schema.Column{WorkflowsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "workflows_organizations_workflows",
				Columns:    []*schema.Column{WorkflowsColumns[9]},
				RefColumns: []*schema.Column{OrganizationsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "workflows_workflow_contracts_contract",
				Columns:    []*schema.Column{WorkflowsColumns[10]},
				RefColumns: []*schema.Column{WorkflowContractsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "workflow_name_organization_id",
				Unique:  true,
				Columns: []*schema.Column{WorkflowsColumns[1], WorkflowsColumns[9]},
				Annotation: &entsql.IndexAnnotation{
					Where: "deleted_at IS NULL",
				},
			},
		},
	}
	// WorkflowContractsColumns holds the columns for the "workflow_contracts" table.
	WorkflowContractsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "name", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime, Default: "CURRENT_TIMESTAMP"},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "description", Type: field.TypeString, Nullable: true},
		{Name: "organization_workflow_contracts", Type: field.TypeUUID, Nullable: true},
	}
	// WorkflowContractsTable holds the schema information for the "workflow_contracts" table.
	WorkflowContractsTable = &schema.Table{
		Name:       "workflow_contracts",
		Columns:    WorkflowContractsColumns,
		PrimaryKey: []*schema.Column{WorkflowContractsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "workflow_contracts_organizations_workflow_contracts",
				Columns:    []*schema.Column{WorkflowContractsColumns[5]},
				RefColumns: []*schema.Column{OrganizationsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "workflowcontract_name_organization_workflow_contracts",
				Unique:  true,
				Columns: []*schema.Column{WorkflowContractsColumns[1], WorkflowContractsColumns[5]},
				Annotation: &entsql.IndexAnnotation{
					Where: "deleted_at IS NULL",
				},
			},
		},
	}
	// WorkflowContractVersionsColumns holds the columns for the "workflow_contract_versions" table.
	WorkflowContractVersionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "body", Type: field.TypeBytes},
		{Name: "revision", Type: field.TypeInt, Default: 1},
		{Name: "created_at", Type: field.TypeTime, Default: "CURRENT_TIMESTAMP"},
		{Name: "workflow_contract_versions", Type: field.TypeUUID, Nullable: true},
	}
	// WorkflowContractVersionsTable holds the schema information for the "workflow_contract_versions" table.
	WorkflowContractVersionsTable = &schema.Table{
		Name:       "workflow_contract_versions",
		Columns:    WorkflowContractVersionsColumns,
		PrimaryKey: []*schema.Column{WorkflowContractVersionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "workflow_contract_versions_workflow_contracts_versions",
				Columns:    []*schema.Column{WorkflowContractVersionsColumns[4]},
				RefColumns: []*schema.Column{WorkflowContractsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// WorkflowRunsColumns holds the columns for the "workflow_runs" table.
	WorkflowRunsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeTime, Default: "CURRENT_TIMESTAMP"},
		{Name: "finished_at", Type: field.TypeTime, Nullable: true},
		{Name: "state", Type: field.TypeEnum, Enums: []string{"initialized", "success", "error", "expired", "canceled"}, Default: "initialized"},
		{Name: "reason", Type: field.TypeString, Nullable: true, Size: 2147483647},
		{Name: "run_url", Type: field.TypeString, Nullable: true},
		{Name: "runner_type", Type: field.TypeString, Nullable: true},
		{Name: "attestation", Type: field.TypeJSON, Nullable: true},
		{Name: "attestation_digest", Type: field.TypeString, Nullable: true},
		{Name: "attestation_state", Type: field.TypeBytes, Nullable: true},
		{Name: "contract_revision_used", Type: field.TypeInt},
		{Name: "contract_revision_latest", Type: field.TypeInt},
		{Name: "robot_account_workflowruns", Type: field.TypeUUID, Nullable: true},
		{Name: "workflow_workflowruns", Type: field.TypeUUID, Nullable: true},
		{Name: "workflow_run_contract_version", Type: field.TypeUUID, Nullable: true},
	}
	// WorkflowRunsTable holds the schema information for the "workflow_runs" table.
	WorkflowRunsTable = &schema.Table{
		Name:       "workflow_runs",
		Columns:    WorkflowRunsColumns,
		PrimaryKey: []*schema.Column{WorkflowRunsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "workflow_runs_robot_accounts_workflowruns",
				Columns:    []*schema.Column{WorkflowRunsColumns[12]},
				RefColumns: []*schema.Column{RobotAccountsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "workflow_runs_workflows_workflowruns",
				Columns:    []*schema.Column{WorkflowRunsColumns[13]},
				RefColumns: []*schema.Column{WorkflowsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "workflow_runs_workflow_contract_versions_contract_version",
				Columns:    []*schema.Column{WorkflowRunsColumns[14]},
				RefColumns: []*schema.Column{WorkflowContractVersionsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "workflowrun_created_at_id",
				Unique:  false,
				Columns: []*schema.Column{WorkflowRunsColumns[1], WorkflowRunsColumns[0]},
			},
			{
				Name:    "workflowrun_created_at_state",
				Unique:  false,
				Columns: []*schema.Column{WorkflowRunsColumns[1], WorkflowRunsColumns[3]},
			},
			{
				Name:    "workflowrun_attestation_digest",
				Unique:  false,
				Columns: []*schema.Column{WorkflowRunsColumns[8]},
			},
		},
	}
	// ReferrerReferencesColumns holds the columns for the "referrer_references" table.
	ReferrerReferencesColumns = []*schema.Column{
		{Name: "referrer_id", Type: field.TypeUUID},
		{Name: "referred_by_id", Type: field.TypeUUID},
	}
	// ReferrerReferencesTable holds the schema information for the "referrer_references" table.
	ReferrerReferencesTable = &schema.Table{
		Name:       "referrer_references",
		Columns:    ReferrerReferencesColumns,
		PrimaryKey: []*schema.Column{ReferrerReferencesColumns[0], ReferrerReferencesColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "referrer_references_referrer_id",
				Columns:    []*schema.Column{ReferrerReferencesColumns[0]},
				RefColumns: []*schema.Column{ReferrersColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "referrer_references_referred_by_id",
				Columns:    []*schema.Column{ReferrerReferencesColumns[1]},
				RefColumns: []*schema.Column{ReferrersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// ReferrerWorkflowsColumns holds the columns for the "referrer_workflows" table.
	ReferrerWorkflowsColumns = []*schema.Column{
		{Name: "referrer_id", Type: field.TypeUUID},
		{Name: "workflow_id", Type: field.TypeUUID},
	}
	// ReferrerWorkflowsTable holds the schema information for the "referrer_workflows" table.
	ReferrerWorkflowsTable = &schema.Table{
		Name:       "referrer_workflows",
		Columns:    ReferrerWorkflowsColumns,
		PrimaryKey: []*schema.Column{ReferrerWorkflowsColumns[0], ReferrerWorkflowsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "referrer_workflows_referrer_id",
				Columns:    []*schema.Column{ReferrerWorkflowsColumns[0]},
				RefColumns: []*schema.Column{ReferrersColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "referrer_workflows_workflow_id",
				Columns:    []*schema.Column{ReferrerWorkflowsColumns[1]},
				RefColumns: []*schema.Column{WorkflowsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// WorkflowRunCasBackendsColumns holds the columns for the "workflow_run_cas_backends" table.
	WorkflowRunCasBackendsColumns = []*schema.Column{
		{Name: "workflow_run_id", Type: field.TypeUUID},
		{Name: "cas_backend_id", Type: field.TypeUUID},
	}
	// WorkflowRunCasBackendsTable holds the schema information for the "workflow_run_cas_backends" table.
	WorkflowRunCasBackendsTable = &schema.Table{
		Name:       "workflow_run_cas_backends",
		Columns:    WorkflowRunCasBackendsColumns,
		PrimaryKey: []*schema.Column{WorkflowRunCasBackendsColumns[0], WorkflowRunCasBackendsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "workflow_run_cas_backends_workflow_run_id",
				Columns:    []*schema.Column{WorkflowRunCasBackendsColumns[0]},
				RefColumns: []*schema.Column{WorkflowRunsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "workflow_run_cas_backends_cas_backend_id",
				Columns:    []*schema.Column{WorkflowRunCasBackendsColumns[1]},
				RefColumns: []*schema.Column{CasBackendsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		APITokensTable,
		CasBackendsTable,
		CasMappingsTable,
		IntegrationsTable,
		IntegrationAttachmentsTable,
		MembershipsTable,
		OrgInvitationsTable,
		OrganizationsTable,
		ReferrersTable,
		RobotAccountsTable,
		UsersTable,
		WorkflowsTable,
		WorkflowContractsTable,
		WorkflowContractVersionsTable,
		WorkflowRunsTable,
		ReferrerReferencesTable,
		ReferrerWorkflowsTable,
		WorkflowRunCasBackendsTable,
	}
)

func init() {
	APITokensTable.ForeignKeys[0].RefTable = OrganizationsTable
	CasBackendsTable.ForeignKeys[0].RefTable = OrganizationsTable
	CasMappingsTable.ForeignKeys[0].RefTable = CasBackendsTable
	CasMappingsTable.ForeignKeys[1].RefTable = WorkflowRunsTable
	CasMappingsTable.ForeignKeys[2].RefTable = OrganizationsTable
	IntegrationsTable.ForeignKeys[0].RefTable = OrganizationsTable
	IntegrationAttachmentsTable.ForeignKeys[0].RefTable = IntegrationsTable
	IntegrationAttachmentsTable.ForeignKeys[1].RefTable = WorkflowsTable
	MembershipsTable.ForeignKeys[0].RefTable = OrganizationsTable
	MembershipsTable.ForeignKeys[1].RefTable = UsersTable
	OrgInvitationsTable.ForeignKeys[0].RefTable = OrganizationsTable
	OrgInvitationsTable.ForeignKeys[1].RefTable = UsersTable
	RobotAccountsTable.ForeignKeys[0].RefTable = WorkflowsTable
	WorkflowsTable.ForeignKeys[0].RefTable = OrganizationsTable
	WorkflowsTable.ForeignKeys[1].RefTable = WorkflowContractsTable
	WorkflowContractsTable.ForeignKeys[0].RefTable = OrganizationsTable
	WorkflowContractVersionsTable.ForeignKeys[0].RefTable = WorkflowContractsTable
	WorkflowRunsTable.ForeignKeys[0].RefTable = RobotAccountsTable
	WorkflowRunsTable.ForeignKeys[1].RefTable = WorkflowsTable
	WorkflowRunsTable.ForeignKeys[2].RefTable = WorkflowContractVersionsTable
	ReferrerReferencesTable.ForeignKeys[0].RefTable = ReferrersTable
	ReferrerReferencesTable.ForeignKeys[1].RefTable = ReferrersTable
	ReferrerWorkflowsTable.ForeignKeys[0].RefTable = ReferrersTable
	ReferrerWorkflowsTable.ForeignKeys[1].RefTable = WorkflowsTable
	WorkflowRunCasBackendsTable.ForeignKeys[0].RefTable = WorkflowRunsTable
	WorkflowRunCasBackendsTable.ForeignKeys[1].RefTable = CasBackendsTable
}
