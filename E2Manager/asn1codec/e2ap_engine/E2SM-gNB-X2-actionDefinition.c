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
 * From ASN.1 module "E2SM-gNB-X2-IEs"
 * 	found in "../../asnFiles/e2sm-gNB-X2-release-1-v041.asn"
 * 	`asn1c -fcompound-names -fincludes-quoted -fno-include-deps -findirect-choice -gen-PER -no-gen-OER -D.`
 */

#include "E2SM-gNB-X2-actionDefinition.h"

#include "ActionParameter-Item.h"
static int
memb_actionParameter_List_constraint_1(const asn_TYPE_descriptor_t *td, const void *sptr,
			asn_app_constraint_failed_f *ctfailcb, void *app_key) {
	size_t size;
	
	if(!sptr) {
		ASN__CTFAIL(app_key, td, sptr,
			"%s: value not given (%s:%d)",
			td->name, __FILE__, __LINE__);
		return -1;
	}
	
	/* Determine the number of elements */
	size = _A_CSEQUENCE_FROM_VOID(sptr)->count;
	
	if((size >= 1 && size <= 255)) {
		/* Perform validation of the inner elements */
		return td->encoding_constraints.general_constraints(td, sptr, ctfailcb, app_key);
	} else {
		ASN__CTFAIL(app_key, td, sptr,
			"%s: constraint failed (%s:%d)",
			td->name, __FILE__, __LINE__);
		return -1;
	}
}

static asn_per_constraints_t asn_PER_type_actionParameter_List_constr_3 CC_NOTUSED = {
	{ APC_UNCONSTRAINED,	-1, -1,  0,  0 },
	{ APC_CONSTRAINED,	 8,  8,  1,  255 }	/* (SIZE(1..255)) */,
	0, 0	/* No PER value map */
};
static asn_per_constraints_t asn_PER_memb_actionParameter_List_constr_3 CC_NOTUSED = {
	{ APC_UNCONSTRAINED,	-1, -1,  0,  0 },
	{ APC_CONSTRAINED,	 8,  8,  1,  255 }	/* (SIZE(1..255)) */,
	0, 0	/* No PER value map */
};
static asn_TYPE_member_t asn_MBR_actionParameter_List_3[] = {
	{ ATF_POINTER, 0, 0,
		(ASN_TAG_CLASS_UNIVERSAL | (16 << 2)),
		0,
		&asn_DEF_ActionParameter_Item,
		0,
		{ 0, 0, 0 },
		0, 0, /* No default value */
		""
		},
};
static const ber_tlv_tag_t asn_DEF_actionParameter_List_tags_3[] = {
	(ASN_TAG_CLASS_CONTEXT | (1 << 2)),
	(ASN_TAG_CLASS_UNIVERSAL | (16 << 2))
};
static asn_SET_OF_specifics_t asn_SPC_actionParameter_List_specs_3 = {
	sizeof(struct E2SM_gNB_X2_actionDefinition__actionParameter_List),
	offsetof(struct E2SM_gNB_X2_actionDefinition__actionParameter_List, _asn_ctx),
	0,	/* XER encoding is XMLDelimitedItemList */
};
static /* Use -fall-defs-global to expose */
asn_TYPE_descriptor_t asn_DEF_actionParameter_List_3 = {
	"actionParameter-List",
	"actionParameter-List",
	&asn_OP_SEQUENCE_OF,
	asn_DEF_actionParameter_List_tags_3,
	sizeof(asn_DEF_actionParameter_List_tags_3)
		/sizeof(asn_DEF_actionParameter_List_tags_3[0]) - 1, /* 1 */
	asn_DEF_actionParameter_List_tags_3,	/* Same as above */
	sizeof(asn_DEF_actionParameter_List_tags_3)
		/sizeof(asn_DEF_actionParameter_List_tags_3[0]), /* 2 */
	{ 0, &asn_PER_type_actionParameter_List_constr_3, SEQUENCE_OF_constraint },
	asn_MBR_actionParameter_List_3,
	1,	/* Single element */
	&asn_SPC_actionParameter_List_specs_3	/* Additional specs */
};

static asn_TYPE_member_t asn_MBR_E2SM_gNB_X2_actionDefinition_1[] = {
	{ ATF_NOFLAGS, 0, offsetof(struct E2SM_gNB_X2_actionDefinition, style_ID),
		(ASN_TAG_CLASS_CONTEXT | (0 << 2)),
		-1,	/* IMPLICIT tag at current level */
		&asn_DEF_Style_ID,
		0,
		{ 0, 0, 0 },
		0, 0, /* No default value */
		"style-ID"
		},
	{ ATF_POINTER, 1, offsetof(struct E2SM_gNB_X2_actionDefinition, actionParameter_List),
		(ASN_TAG_CLASS_CONTEXT | (1 << 2)),
		0,
		&asn_DEF_actionParameter_List_3,
		0,
		{ 0, &asn_PER_memb_actionParameter_List_constr_3,  memb_actionParameter_List_constraint_1 },
		0, 0, /* No default value */
		"actionParameter-List"
		},
};
static const int asn_MAP_E2SM_gNB_X2_actionDefinition_oms_1[] = { 1 };
static const ber_tlv_tag_t asn_DEF_E2SM_gNB_X2_actionDefinition_tags_1[] = {
	(ASN_TAG_CLASS_UNIVERSAL | (16 << 2))
};
static const asn_TYPE_tag2member_t asn_MAP_E2SM_gNB_X2_actionDefinition_tag2el_1[] = {
    { (ASN_TAG_CLASS_CONTEXT | (0 << 2)), 0, 0, 0 }, /* style-ID */
    { (ASN_TAG_CLASS_CONTEXT | (1 << 2)), 1, 0, 0 } /* actionParameter-List */
};
static asn_SEQUENCE_specifics_t asn_SPC_E2SM_gNB_X2_actionDefinition_specs_1 = {
	sizeof(struct E2SM_gNB_X2_actionDefinition),
	offsetof(struct E2SM_gNB_X2_actionDefinition, _asn_ctx),
	asn_MAP_E2SM_gNB_X2_actionDefinition_tag2el_1,
	2,	/* Count of tags in the map */
	asn_MAP_E2SM_gNB_X2_actionDefinition_oms_1,	/* Optional members */
	1, 0,	/* Root/Additions */
	2,	/* First extension addition */
};
asn_TYPE_descriptor_t asn_DEF_E2SM_gNB_X2_actionDefinition = {
	"E2SM-gNB-X2-actionDefinition",
	"E2SM-gNB-X2-actionDefinition",
	&asn_OP_SEQUENCE,
	asn_DEF_E2SM_gNB_X2_actionDefinition_tags_1,
	sizeof(asn_DEF_E2SM_gNB_X2_actionDefinition_tags_1)
		/sizeof(asn_DEF_E2SM_gNB_X2_actionDefinition_tags_1[0]), /* 1 */
	asn_DEF_E2SM_gNB_X2_actionDefinition_tags_1,	/* Same as above */
	sizeof(asn_DEF_E2SM_gNB_X2_actionDefinition_tags_1)
		/sizeof(asn_DEF_E2SM_gNB_X2_actionDefinition_tags_1[0]), /* 1 */
	{ 0, 0, SEQUENCE_constraint },
	asn_MBR_E2SM_gNB_X2_actionDefinition_1,
	2,	/* Elements count */
	&asn_SPC_E2SM_gNB_X2_actionDefinition_specs_1	/* Additional specs */
};

