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
 * From ASN.1 module "E2AP-IEs"
 * 	found in "../../asnFiles/e2ap-v031.asn"
 * 	`asn1c -fcompound-names -fincludes-quoted -fno-include-deps -findirect-choice -gen-PER -no-gen-OER -D.`
 */

#ifndef	_RICcontrolAckRequest_H_
#define	_RICcontrolAckRequest_H_


#include "asn_application.h"

/* Including external dependencies */
#include "NativeEnumerated.h"

#ifdef __cplusplus
extern "C" {
#endif

/* Dependencies */
typedef enum RICcontrolAckRequest {
	RICcontrolAckRequest_noAck	= 0,
	RICcontrolAckRequest_ack	= 1,
	RICcontrolAckRequest_nAck	= 2
	/*
	 * Enumeration is extensible
	 */
} e_RICcontrolAckRequest;

/* RICcontrolAckRequest */
typedef long	 RICcontrolAckRequest_t;

/* Implementation */
extern asn_TYPE_descriptor_t asn_DEF_RICcontrolAckRequest;
asn_struct_free_f RICcontrolAckRequest_free;
asn_struct_print_f RICcontrolAckRequest_print;
asn_constr_check_f RICcontrolAckRequest_constraint;
ber_type_decoder_f RICcontrolAckRequest_decode_ber;
der_type_encoder_f RICcontrolAckRequest_encode_der;
xer_type_decoder_f RICcontrolAckRequest_decode_xer;
xer_type_encoder_f RICcontrolAckRequest_encode_xer;
per_type_decoder_f RICcontrolAckRequest_decode_uper;
per_type_encoder_f RICcontrolAckRequest_encode_uper;
per_type_decoder_f RICcontrolAckRequest_decode_aper;
per_type_encoder_f RICcontrolAckRequest_encode_aper;

#ifdef __cplusplus
}
#endif

#endif	/* _RICcontrolAckRequest_H_ */
#include "asn_internal.h"
