// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package wafv2

// **PLEASE DELETE THIS AND ALL TIP COMMENTS BEFORE SUBMITTING A PR FOR REVIEW!**
//
// TIP: ==== INTRODUCTION ====
// Thank you for trying the skaff tool!
//
// You have opted to include these helpful comments. They all include "TIP:"
// to help you find and remove them when you're done with them.
//
// While some aspects of this file are customized to your input, the
// scaffold tool does *not* look at the AWS API and ensure it has correct
// function, structure, and variable names. It makes guesses based on
// commonalities. You will need to make significant adjustments.
//
// In other words, as generated, this is a rough outline of the work you will
// need to do. If something doesn't make sense for your situation, get rid of
// it.

import (
	// TIP: ==== IMPORTS ====
	// This is a common set of imports but not customized to your code since
	// your code hasn't been written yet. Make sure you, your IDE, or
	// goimports -w <file> fixes these imports.
	//
	// The provider linter wants your imports to be in two groups: first,
	// standard library (i.e., "fmt" or "strings"), second, everything else.
	//
	// Also, AWS Go SDK v2 may handle nested structures differently than v1,
	// using the services/wafv2/types package. If so, you'll
	// need to import types and reference the nested types, e.g., as
	// awstypes.<Type Name>.
	"context"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/wafv2"
	awstypes "github.com/aws/aws-sdk-go-v2/service/wafv2/types"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
    "github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-aws/internal/create"
	"github.com/hashicorp/terraform-provider-aws/internal/framework"
	"github.com/hashicorp/terraform-provider-aws/internal/framework/flex"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	 fwtypes "github.com/hashicorp/terraform-provider-aws/internal/framework/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

// TIP: ==== FILE STRUCTURE ====
// All data sources should follow this basic outline. Improve this data source's
// maintainability by sticking to it.
//
// 1. Package declaration
// 2. Imports
// 3. Main data source struct with schema method
// 4. Read method
// 5. Other functions (flatteners, expanders, waiters, finders, etc.)

// Function annotations are used for datasource registration to the Provider. DO NOT EDIT.
// @FrameworkDataSource("aws_wafv2_managed_rule_group", name="Managed Rule Group")
func newDataSourceManagedRuleGroup(context.Context) (datasource.DataSourceWithConfigure, error) {
	return &dataSourceManagedRuleGroup{}, nil
}

const (
	DSNameManagedRuleGroup = "Managed Rule Group Data Source"
)

type dataSourceManagedRuleGroup struct {
	framework.DataSourceWithConfigure
}

func (d *dataSourceManagedRuleGroup) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) { // nosemgrep:ci.meta-in-func-name
	resp.TypeName = "aws_wafv2_managed_rule_group"
}

// TIP: ==== SCHEMA ====
// In the schema, add each of the arguments and attributes in snake
// case (e.g., delete_automated_backups).
// * Alphabetize arguments to make them easier to find.
// * Do not add a blank line between arguments/attributes.
//
// Users can configure argument values while attribute values cannot be
// configured and are used as output. Arguments have either:
// Required: true,
// Optional: true,
//
// All attributes will be computed and some arguments. If users will
// want to read updated information or detect drift for an argument,
// it should be computed:
// Computed: true,
//
// You will typically find arguments in the input struct
// (e.g., CreateDBInstanceInput) for the create operation. Sometimes
// they are only in the input struct (e.g., ModifyDBInstanceInput) for
// the modify operation.
//
// For more about schema options, visit
// https://developer.hashicorp.com/terraform/plugin/framework/handling-data/schemas?page=schemas
func (d *dataSourceManagedRuleGroup) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Data for Vendor Maintained ManagedRuleGroups",
		Attributes: map[string]schema.Attribute{
			"vendor_name": schema.StringAttribute{
				Required: true,
			},
			names.AttrName: schema.StringAttribute{
				Required: true,
			},
			names.AttrScope: schema.StringAttribute{
				Required: true,
				Validators: []validator.String{
					stringvalidator.OneOf(
						string(awstypes.ScopeCloudfront),
						string(awstypes.ScopeRegional),
					),
				},
			},
			names.AttrSNSTopicARN: framework.ARNAttributeComputedOnly(),
			// names.AttrDescription: schema.StringAttribute{
			// 	Computed: true,
			// },
			names.AttrVersion: schema.StringAttribute{
				Optional: true,
			},
			"capacity": schema.NumberAttribute{
				Computed: true,
			},
			// "versioning_supported": schema.BoolAttribute{
			// 	Computed: true,
			// },
			"rules": schema.ListAttribute{
				CustomType: fwtypes.NewListNestedObjectTypeOf[RuleSummaryModel](ctx),
				ElementType: fwtypes.NewObjectTypeOf[RuleSummaryModel](ctx),
				Computed: true,
			},
			"available_labels": schema.ListAttribute{
				CustomType: fwtypes.NewListNestedObjectTypeOf[LabelModel](ctx),
				ElementType: fwtypes.NewObjectTypeOf[LabelModel](ctx),
				Computed: true,
			},
			"versions": schema.ListAttribute{
				CustomType: fwtypes.NewListNestedObjectTypeOf[ManagedRuleGroupVersionModel](ctx),
				ElementType: fwtypes.NewObjectTypeOf[ManagedRuleGroupVersionModel](ctx),
				Computed: true,
			},
			"current_default_version": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}
// TIP: ==== ASSIGN CRUD METHODS ====
// Data sources only have a read method.
func (d *dataSourceManagedRuleGroup) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	// TIP: ==== DATA SOURCE READ ====
	// Generally, the Read function should do the following things. Make
	// sure there is a good reason if you don't do one of these.
	//
	// 1. Get a client connection to the relevant service
	// 2. Fetch the config
	// 3. Get information about a resource from AWS
	// 4. Set the ID, arguments, and attributes
	// 5. Set the tags
	// 6. Set the state
	// TIP: -- 1. Get a client connection to the relevant service
	conn := d.Meta().WAFV2Client(ctx)

	// TIP: -- 2. Fetch the config
	var data dataSourceManagedRuleGroupModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// TIP: -- 3. Get information about a resource from AWS
	version := data.Version.ValueString()
	if data.Version.IsNull() {
		version = ""
	}
	out, err := describeManagedRuleGroup(ctx, conn, data.Name.ValueString(), data.Scope.ValueString(), data.VendorName.ValueString(), version)
	if err != nil {
		resp.Diagnostics.AddError(
			create.ProblemStandardMessage(names.WAFV2, 
				create.ErrActionReading, 
				DSNameManagedRuleGroup, 
				strings.Join([]string{data.Scope.ValueString(), data.VendorName.ValueString(), data.Name.ValueString(), version},":"), 
				err),
			err.Error(),
		)
		return
	}

	// TIP: -- 4. Set the ID, arguments, and attributes
	// Using a field name prefix allows mapping fields such as `ManagedRuleGroupId` to `ID`
	resp.Diagnostics.Append(flex.Flatten(ctx, out, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	
	outVer, err := listAvailableManagedRuleGroupVersions(ctx, conn, data.Name.ValueString(), data.Scope.ValueString(), data.VendorName.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			create.ProblemStandardMessage(names.WAFV2, 
				create.ErrActionReading, 
				DSNameManagedRuleGroup, 
				strings.Join([]string{data.Scope.ValueString(), data.VendorName.ValueString(), data.Name.ValueString()},":"), 
				err),
			err.Error(),
		)
		return
	}

	resp.Diagnostics.Append(flex.Flatten(ctx, outVer, &data)...)

	// TIP: -- 5. Set the tags

	// TIP: -- 6. Set the state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}


// TIP: ==== DATA STRUCTURES ====
// With Terraform Plugin-Framework configurations are deserialized into
// Go types, providing type safety without the need for type assertions.
// These structs should match the schema definition exactly, and the `tfsdk`
// tag value should match the attribute name.
//
// Nested objects are represented in their own data struct. These will
// also have a corresponding attribute type mapping for use inside flex
// functions.
//
// See more:
// https://developer.hashicorp.com/terraform/plugin/framework/handling-data/accessing-values
type dataSourceManagedRuleGroupModel struct {
	VendorName				types.String														`tfsdk:"vendor_name"`
	Name 					types.String														`tfsdk:"name"`
	Scope					types.String														`tfsdk:"scope"`
	SNSTopicArn 			types.String														`tfsdk:"sns_topic_arn"`
	// Description 			types.String														`tfsdk:"description"`
	Version 				types.String														`tfsdk:"version"`
	Capactity   			types.Number														`tfsdk:"capacity"`
	// VersioningSupported		types.Bool															`tfsdk:"versioning_supported"`
	Rules               	fwtypes.ListNestedObjectValueOf[RuleSummaryModel]				`tfsdk:"rules"`
	AvailableLabels			fwtypes.ListNestedObjectValueOf[LabelModel]							`tfsdk:"available_labels"`
	Versions				fwtypes.ListNestedObjectValueOf[ManagedRuleGroupVersionModel]			`tfsdk:"versions"`
	CurrentDefaultVersion	types.String														`tfsdk:"current_default_version"`
}

func describeManagedRuleGroup(ctx context.Context, conn *wafv2.Client, name string, scope string, vendor string, version string) (*wafv2.DescribeManagedRuleGroupOutput, error){
	input := &wafv2.DescribeManagedRuleGroupInput{
		Name: aws.String(name),
		VendorName: aws.String(vendor),
		Scope: awstypes.Scope(scope),
	}

	if version != "" {
		input.VersionName = aws.String(version)
	}

	return conn.DescribeManagedRuleGroup(ctx, input)
}

func listAvailableManagedRuleGroupVersions(ctx context.Context, conn *wafv2.Client, name string, scope string, vendor string) (*wafv2.ListAvailableManagedRuleGroupVersionsOutput, error){
	input := &wafv2.ListAvailableManagedRuleGroupVersionsInput{
		Name: aws.String(name),
		VendorName: aws.String(vendor),
		Scope: awstypes.Scope(scope),
	}

	return conn.ListAvailableManagedRuleGroupVersions(ctx, input)
}

type ManagedRuleGroupVersionModel struct{
	Name				types.String		`tfsdk:"name"`
	LastUpdateTimeStamp types.String		`tfsdk:"last_update_timestamp"`
}

func (m *ManagedRuleGroupVersionModel) Flatten(ctx context.Context, v any) diag.Diagnostics {
    var diags diag.Diagnostics
    
    version := v.(awstypes.ManagedRuleGroupVersion)
    m.Name = types.StringPointerValue(version.Name)

    if version.LastUpdateTimestamp != nil {
        // Convert time.Time to RFC3339 string format
        m.LastUpdateTimeStamp = types.StringValue(version.LastUpdateTimestamp.Format(time.RFC3339))
    } else {
        m.LastUpdateTimeStamp = types.StringNull()
    }

    return diags
}


type LabelModel struct{
	Name				types.String		`tfsdk:"name"`
}

type RuleSummaryModel struct{
	Name	types.String	`tfsdk:"name"`
	// only care about action type
	Action	types.String	`tfsdk:"action"`
}

func (m *RuleSummaryModel) Flatten(ctx context.Context, v any) diag.Diagnostics {
    var diags diag.Diagnostics
    
    rs := v.(awstypes.RuleSummary)
    m.Name = types.StringPointerValue(rs.Name)

    // Determine action type based on which field is present
    if rs.Action != nil {
        switch {
        case rs.Action.Block != nil:
            m.Action = types.StringValue("BLOCK")
        case rs.Action.Allow != nil:
            m.Action = types.StringValue("ALLOW")
        case rs.Action.Count != nil:
            m.Action = types.StringValue("COUNT")
        case rs.Action.Captcha != nil:
            m.Action = types.StringValue("CAPTCHA")
        case rs.Action.Challenge != nil:
            m.Action = types.StringValue("CHALLENGE")
        }
    } else {
        m.Action = types.StringNull()
    }

    return diags
}
