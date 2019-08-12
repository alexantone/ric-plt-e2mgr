/*
 *
 * Copyright 2019 AT&T Intellectual Property
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */


/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "X2AP-PDU-Contents"
 * 	found in "../../asnFiles/X2AP-PDU-Contents.asn"
 * 	`asn1c -fcompound-names -fincludes-quoted -fno-include-deps -findirect-choice -gen-PER -no-gen-OER -D.`
 */

#include "ServedNRCellsToModify-Item.h"

#include "NRNeighbour-Information.h"
#include "ProtocolExtensionContainer.h"
asn_TYPE_member_t asn_MBR_ServedNRCellsToModify_Item_1[] = {
	{ ATF_NOFLAGS, 0, offsetof(struct ServedNRCellsToModify_Item, old_nrcgi),
		(ASN_TAG_CLASS_CONTEXT | (0 << 2)),
		-1,	/* IMPLICIT tag at current level */
		&asn_DEF_NRCGI,
		0,
		{ 0, 0, 0 },
		0, 0, /* No default value */
		"old-nrcgi"
		},
	{ ATF_NOFLAGS, 0, offsetof(struct ServedNRCellsToModify_Item, servedNRCellInformation),
		(ASN_TAG_CLASS_CONTEXT | (1 << 2)),
		-1,	/* IMPLICIT tag at current level */
		&asn_DEF_ServedNRCell_Information,
		0,
		{ 0, 0, 0 },
		0, 0, /* No default value */
		"servedNRCellInformation"
		},
	{ ATF_POINTER, 3, offsetof(struct ServedNRCellsToModify_Item, nrNeighbourInformation),
		(ASN_TAG_CLASS_CONTEXT | (2 << 2)),
		-1,	/* IMPLICIT tag at current level */
		&asn_DEF_NRNeighbour_Information,
		0,
		{ 0, 0, 0 },
		0, 0, /* No default value */
		"nrNeighbourInformation"
		},
	{ ATF_POINTER, 2, offsetof(struct ServedNRCellsToModify_Item, nrDeactivationIndication),
		(ASN_TAG_CLASS_CONTEXT | (3 << 2)),
		-1,	/* IMPLICIT tag at current level */
		&asn_DEF_DeactivationIndication,
		0,
		{ 0, 0, 0 },
		0, 0, /* No default value */
		"nrDeactivationIndication"
		},
	{ ATF_POINTER, 1, offsetof(struct ServedNRCellsToModify_Item, iE_Extensions),
		(ASN_TAG_CLASS_CONTEXT | (4 << 2)),
		-1,	/* IMPLICIT tag at current level */
		&asn_DEF_ProtocolExtensionContainer_170P100,
		0,
		{ 0, 0, 0 },
		0, 0, /* No default value */
		"iE-Extensions"
		},
};
static const int asn_MAP_ServedNRCellsToModify_Item_oms_1[] = { 2, 3, 4 };
static const ber_tlv_tag_t asn_DEF_ServedNRCellsToModify_Item_tags_1[] = {
	(ASN_TAG_CLASS_UNIVERSAL | (16 << 2))
};
static const asn_TYPE_tag2member_t asn_MAP_ServedNRCellsToModify_Item_tag2el_1[] = {
    { (ASN_TAG_CLASS_CONTEXT | (0 << 2)), 0, 0, 0 }, /* old-nrcgi */
    { (ASN_TAG_CLASS_CONTEXT | (1 << 2)), 1, 0, 0 }, /* servedNRCellInformation */
    { (ASN_TAG_CLASS_CONTEXT | (2 << 2)), 2, 0, 0 }, /* nrNeighbourInformation */
    { (ASN_TAG_CLASS_CONTEXT | (3 << 2)), 3, 0, 0 }, /* nrDeactivationIndication */
    { (ASN_TAG_CLASS_CONTEXT | (4 << 2)), 4, 0, 0 } /* iE-Extensions */
};
asn_SEQUENCE_specifics_t asn_SPC_ServedNRCellsToModify_Item_specs_1 = {
	sizeof(struct ServedNRCellsToModify_Item),
	offsetof(struct ServedNRCellsToModify_Item, _asn_ctx),
	asn_MAP_ServedNRCellsToModify_Item_tag2el_1,
	5,	/* Count of tags in the map */
	asn_MAP_ServedNRCellsToModify_Item_oms_1,	/* Optional members */
	3, 0,	/* Root/Additions */
	5,	/* First extension addition */
};
asn_TYPE_descriptor_t asn_DEF_ServedNRCellsToModify_Item = {
	"ServedNRCellsToModify-Item",
	"ServedNRCellsToModify-Item",
	&asn_OP_SEQUENCE,
	asn_DEF_ServedNRCellsToModify_Item_tags_1,
	sizeof(asn_DEF_ServedNRCellsToModify_Item_tags_1)
		/sizeof(asn_DEF_ServedNRCellsToModify_Item_tags_1[0]), /* 1 */
	asn_DEF_ServedNRCellsToModify_Item_tags_1,	/* Same as above */
	sizeof(asn_DEF_ServedNRCellsToModify_Item_tags_1)
		/sizeof(asn_DEF_ServedNRCellsToModify_Item_tags_1[0]), /* 1 */
	{ 0, 0, SEQUENCE_constraint },
	asn_MBR_ServedNRCellsToModify_Item_1,
	5,	/* Elements count */
	&asn_SPC_ServedNRCellsToModify_Item_specs_1	/* Additional specs */
};

