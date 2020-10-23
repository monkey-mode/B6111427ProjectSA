/* tslint:disable */
/* eslint-disable */
/**
 * SUT SA Example API
 * This is a sample server for SUT SE 2563
 *
 * The version of the OpenAPI document: 1.0
 * Contact: support@swagger.io
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import {
    EntStatusEdges,
    EntStatusEdgesFromJSON,
    EntStatusEdgesFromJSONTyped,
    EntStatusEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntStatus
 */
export interface EntStatus {
    /**
     * STATUSNAME holds the value of the "STATUS_NAME" field.
     * @type {string}
     * @memberof EntStatus
     */
    sTATUSNAME?: string;
    /**
     * 
     * @type {EntStatusEdges}
     * @memberof EntStatus
     */
    edges?: EntStatusEdges;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntStatus
     */
    id?: number;
}

export function EntStatusFromJSON(json: any): EntStatus {
    return EntStatusFromJSONTyped(json, false);
}

export function EntStatusFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntStatus {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'sTATUSNAME': !exists(json, 'STATUS_NAME') ? undefined : json['STATUS_NAME'],
        'edges': !exists(json, 'edges') ? undefined : EntStatusEdgesFromJSON(json['edges']),
        'id': !exists(json, 'id') ? undefined : json['id'],
    };
}

export function EntStatusToJSON(value?: EntStatus | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'STATUS_NAME': value.sTATUSNAME,
        'edges': EntStatusEdgesToJSON(value.edges),
        'id': value.id,
    };
}


