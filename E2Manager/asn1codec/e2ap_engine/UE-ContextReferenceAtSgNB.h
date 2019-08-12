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

#ifndef	_UE_ContextReferenceAtSgNB_H_
#define	_UE_ContextReferenceAtSgNB_H_


#include "asn_application.h"

/* Including external dependencies */
#include "GlobalGNB-ID.h"
#include "SgNB-UE-X2AP-ID.h"
#include "constr_SEQUENCE.h"

#ifdef __cplusplus
extern "C" {
#endif

/* Forward declarations */
struct ProtocolExtensionContainer;

/* UE-ContextReferenceAtSgNB */
typedef struct UE_ContextReferenceAtSgNB {
	GlobalGNB_ID_t	 source_GlobalSgNB_ID;
	SgNB_UE_X2AP_ID_t	 sgNB_UE_X2AP_ID;
	struct ProtocolExtensionContainer	*iE_Extensions;	/* OPTIONAL */
	/*
	 * This type is extensible,
	 * possible extensions are below.
	 */
	
	/* Context for parsing across buffer boundaries */
	asn_struct_ctx_t _asn_ctx;
} UE_ContextReferenceAtSgNB_t;

/* Implementation */
extern asn_TYPE_descriptor_t asn_DEF_UE_ContextReferenceAtSgNB;

#ifdef __cplusplus
}
#endif

#endif	/* _UE_ContextReferenceAtSgNB_H_ */
#include "asn_internal.h"
