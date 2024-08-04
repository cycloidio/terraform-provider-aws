// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package directconnect

// Exports for use in tests only.
var (
	ResourceBGPPeer                    = resourceBGPPeer
	ResourceConnection                 = resourceConnection
	ResourceConnectionAssociation      = resourceConnectionAssociation
	ResourceConnectionConfirmation     = resourceConnectionConfirmation
	ResourceGateway                    = resourceGateway
	ResourceGatewayAssociation         = resourceGatewayAssociation
	ResourceGatewayAssociationProposal = resourceGatewayAssociationProposal

	FindBGPPeerByThreePartKey          = findBGPPeerByThreePartKey
	FindConnectionByID                 = findConnectionByID
	FindConnectionLAGAssociation       = findConnectionLAGAssociation
	FindGatewayAssociationByID         = findGatewayAssociationByID
	FindGatewayAssociationProposalByID = findGatewayAssociationProposalByID
	FindGatewayByID                    = findGatewayByID
	FindVirtualInterfaceByID           = findVirtualInterfaceByID
	GatewayAssociationStateUpgradeV0   = gatewayAssociationStateUpgradeV0
	ValidConnectionBandWidth           = validConnectionBandWidth
)
