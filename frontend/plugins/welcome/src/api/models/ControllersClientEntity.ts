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
/**
 * 
 * @export
 * @interface ControllersClientEntity
 */
export interface ControllersClientEntity {
    /**
     * 
     * @type {number}
     * @memberof ControllersClientEntity
     */
    sid?: number;
}

export function ControllersClientEntityFromJSON(json: any): ControllersClientEntity {
    return ControllersClientEntityFromJSONTyped(json, false);
}

export function ControllersClientEntityFromJSONTyped(json: any, ignoreDiscriminator: boolean): ControllersClientEntity {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'sid': !exists(json, 'sid') ? undefined : json['sid'],
    };
}

export function ControllersClientEntityToJSON(value?: ControllersClientEntity | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'sid': value.sid,
    };
}


