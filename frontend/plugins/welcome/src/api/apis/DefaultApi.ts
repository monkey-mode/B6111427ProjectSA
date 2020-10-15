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


import * as runtime from '../runtime';
import {
    ControllersBooking,
    ControllersBookingFromJSON,
    ControllersBookingToJSON,
    EntBooking,
    EntBookingFromJSON,
    EntBookingToJSON,
    EntBookingtype,
    EntBookingtypeFromJSON,
    EntBookingtypeToJSON,
    EntClientEntity,
    EntClientEntityFromJSON,
    EntClientEntityToJSON,
    EntRole,
    EntRoleFromJSON,
    EntRoleToJSON,
    EntUser,
    EntUserFromJSON,
    EntUserToJSON,
} from '../models';

export interface CreateBookingRequest {
    booking: ControllersBooking;
}

export interface CreateBookingtypeRequest {
    bookingtype: EntBookingtype;
}

export interface CreateCliententityRequest {
    cliententity: EntClientEntity;
}

export interface CreateRoleRequest {
    role: EntRole;
}

export interface CreateUserRequest {
    user: EntUser;
}

export interface DeleteBookingRequest {
    id: number;
}

export interface DeleteBookingtypeRequest {
    id: number;
}

export interface DeleteCliententityRequest {
    id: number;
}

export interface DeleteRoleRequest {
    id: number;
}

export interface DeleteUserRequest {
    id: number;
}

export interface GetBookingRequest {
    id: number;
}

export interface GetBookingtypeRequest {
    id: number;
}

export interface GetCliententityRequest {
    id: number;
}

export interface GetRoleRequest {
    id: number;
}

export interface GetUserRequest {
    id: number;
}

export interface ListBookingRequest {
    limit?: number;
    offset?: number;
}

export interface ListBookingtypeRequest {
    limit?: number;
    offset?: number;
}

export interface ListCliententityRequest {
    limit?: number;
    offset?: number;
    offset2?: number;
}

export interface ListRoleRequest {
    limit?: number;
    offset?: number;
}

export interface ListUserRequest {
    limit?: number;
    offset?: number;
}

export interface UpdateBookingRequest {
    id: number;
    booking: EntBooking;
}

export interface UpdateBookingtypeRequest {
    id: number;
    bookingtype: EntBookingtype;
}

export interface UpdateCliententityRequest {
    id: number;
    cliententity: EntClientEntity;
}

export interface UpdateRoleRequest {
    id: number;
    role: EntRole;
}

export interface UpdateUserRequest {
    id: number;
    user: EntUser;
}

/**
 * 
 */
export class DefaultApi extends runtime.BaseAPI {

    /**
     * Create booking
     * Create booking
     */
    async createBookingRaw(requestParameters: CreateBookingRequest): Promise<runtime.ApiResponse<EntBooking>> {
        if (requestParameters.booking === null || requestParameters.booking === undefined) {
            throw new runtime.RequiredError('booking','Required parameter requestParameters.booking was null or undefined when calling createBooking.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/bookings`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: ControllersBookingToJSON(requestParameters.booking),
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => EntBookingFromJSON(jsonValue));
    }

    /**
     * Create booking
     * Create booking
     */
    async createBooking(requestParameters: CreateBookingRequest): Promise<EntBooking> {
        const response = await this.createBookingRaw(requestParameters);
        return await response.value();
    }

    /**
     * Create bookingtype
     * Create bookingtype
     */
    async createBookingtypeRaw(requestParameters: CreateBookingtypeRequest): Promise<runtime.ApiResponse<EntBookingtype>> {
        if (requestParameters.bookingtype === null || requestParameters.bookingtype === undefined) {
            throw new runtime.RequiredError('bookingtype','Required parameter requestParameters.bookingtype was null or undefined when calling createBookingtype.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/bookingtypes`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: EntBookingtypeToJSON(requestParameters.bookingtype),
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => EntBookingtypeFromJSON(jsonValue));
    }

    /**
     * Create bookingtype
     * Create bookingtype
     */
    async createBookingtype(requestParameters: CreateBookingtypeRequest): Promise<EntBookingtype> {
        const response = await this.createBookingtypeRaw(requestParameters);
        return await response.value();
    }

    /**
     * Create cliententity
     * Create cliententity
     */
    async createCliententityRaw(requestParameters: CreateCliententityRequest): Promise<runtime.ApiResponse<EntClientEntity>> {
        if (requestParameters.cliententity === null || requestParameters.cliententity === undefined) {
            throw new runtime.RequiredError('cliententity','Required parameter requestParameters.cliententity was null or undefined when calling createCliententity.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/cliententitys`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: EntClientEntityToJSON(requestParameters.cliententity),
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => EntClientEntityFromJSON(jsonValue));
    }

    /**
     * Create cliententity
     * Create cliententity
     */
    async createCliententity(requestParameters: CreateCliententityRequest): Promise<EntClientEntity> {
        const response = await this.createCliententityRaw(requestParameters);
        return await response.value();
    }

    /**
     * Create role
     * Create role
     */
    async createRoleRaw(requestParameters: CreateRoleRequest): Promise<runtime.ApiResponse<EntRole>> {
        if (requestParameters.role === null || requestParameters.role === undefined) {
            throw new runtime.RequiredError('role','Required parameter requestParameters.role was null or undefined when calling createRole.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/roles`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: EntRoleToJSON(requestParameters.role),
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => EntRoleFromJSON(jsonValue));
    }

    /**
     * Create role
     * Create role
     */
    async createRole(requestParameters: CreateRoleRequest): Promise<EntRole> {
        const response = await this.createRoleRaw(requestParameters);
        return await response.value();
    }

    /**
     * Create user
     * Create user
     */
    async createUserRaw(requestParameters: CreateUserRequest): Promise<runtime.ApiResponse<EntUser>> {
        if (requestParameters.user === null || requestParameters.user === undefined) {
            throw new runtime.RequiredError('user','Required parameter requestParameters.user was null or undefined when calling createUser.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/users`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: EntUserToJSON(requestParameters.user),
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => EntUserFromJSON(jsonValue));
    }

    /**
     * Create user
     * Create user
     */
    async createUser(requestParameters: CreateUserRequest): Promise<EntUser> {
        const response = await this.createUserRaw(requestParameters);
        return await response.value();
    }

    /**
     * get booking by ID
     * Delete a booking entity by ID
     */
    async deleteBookingRaw(requestParameters: DeleteBookingRequest): Promise<runtime.ApiResponse<object>> {
        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling deleteBooking.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/bookings/{id}`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'DELETE',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse<any>(response);
    }

    /**
     * get booking by ID
     * Delete a booking entity by ID
     */
    async deleteBooking(requestParameters: DeleteBookingRequest): Promise<object> {
        const response = await this.deleteBookingRaw(requestParameters);
        return await response.value();
    }

    /**
     * get bookingtype by ID
     * Delete a bookingtype entity by ID
     */
    async deleteBookingtypeRaw(requestParameters: DeleteBookingtypeRequest): Promise<runtime.ApiResponse<object>> {
        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling deleteBookingtype.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/bookingtypes/{id}`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'DELETE',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse<any>(response);
    }

    /**
     * get bookingtype by ID
     * Delete a bookingtype entity by ID
     */
    async deleteBookingtype(requestParameters: DeleteBookingtypeRequest): Promise<object> {
        const response = await this.deleteBookingtypeRaw(requestParameters);
        return await response.value();
    }

    /**
     * get cliententity by ID
     * Delete a cliententity entity by ID
     */
    async deleteCliententityRaw(requestParameters: DeleteCliententityRequest): Promise<runtime.ApiResponse<object>> {
        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling deleteCliententity.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/cliententitys/{id}`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'DELETE',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse<any>(response);
    }

    /**
     * get cliententity by ID
     * Delete a cliententity entity by ID
     */
    async deleteCliententity(requestParameters: DeleteCliententityRequest): Promise<object> {
        const response = await this.deleteCliententityRaw(requestParameters);
        return await response.value();
    }

    /**
     * get role by ID
     * Delete a role entity by ID
     */
    async deleteRoleRaw(requestParameters: DeleteRoleRequest): Promise<runtime.ApiResponse<object>> {
        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling deleteRole.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/roles/{id}`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'DELETE',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse<any>(response);
    }

    /**
     * get role by ID
     * Delete a role entity by ID
     */
    async deleteRole(requestParameters: DeleteRoleRequest): Promise<object> {
        const response = await this.deleteRoleRaw(requestParameters);
        return await response.value();
    }

    /**
     * get user by ID
     * Delete a user entity by ID
     */
    async deleteUserRaw(requestParameters: DeleteUserRequest): Promise<runtime.ApiResponse<object>> {
        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling deleteUser.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/users/{id}`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'DELETE',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse<any>(response);
    }

    /**
     * get user by ID
     * Delete a user entity by ID
     */
    async deleteUser(requestParameters: DeleteUserRequest): Promise<object> {
        const response = await this.deleteUserRaw(requestParameters);
        return await response.value();
    }

    /**
     * get booking by ID
     * Get a booking entity by ID
     */
    async getBookingRaw(requestParameters: GetBookingRequest): Promise<runtime.ApiResponse<EntBooking>> {
        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling getBooking.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/bookings/{id}`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => EntBookingFromJSON(jsonValue));
    }

    /**
     * get booking by ID
     * Get a booking entity by ID
     */
    async getBooking(requestParameters: GetBookingRequest): Promise<EntBooking> {
        const response = await this.getBookingRaw(requestParameters);
        return await response.value();
    }

    /**
     * get bookingtype by ID
     * Get a bookingtype entity by ID
     */
    async getBookingtypeRaw(requestParameters: GetBookingtypeRequest): Promise<runtime.ApiResponse<EntBookingtype>> {
        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling getBookingtype.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/bookingtypes/{id}`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => EntBookingtypeFromJSON(jsonValue));
    }

    /**
     * get bookingtype by ID
     * Get a bookingtype entity by ID
     */
    async getBookingtype(requestParameters: GetBookingtypeRequest): Promise<EntBookingtype> {
        const response = await this.getBookingtypeRaw(requestParameters);
        return await response.value();
    }

    /**
     * get cliententity by ID
     * Get a cliententity entity by ID
     */
    async getCliententityRaw(requestParameters: GetCliententityRequest): Promise<runtime.ApiResponse<EntClientEntity>> {
        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling getCliententity.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/cliententitys/{id}`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => EntClientEntityFromJSON(jsonValue));
    }

    /**
     * get cliententity by ID
     * Get a cliententity entity by ID
     */
    async getCliententity(requestParameters: GetCliententityRequest): Promise<EntClientEntity> {
        const response = await this.getCliententityRaw(requestParameters);
        return await response.value();
    }

    /**
     * get role by ID
     * Get a role entity by ID
     */
    async getRoleRaw(requestParameters: GetRoleRequest): Promise<runtime.ApiResponse<EntRole>> {
        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling getRole.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/roles/{id}`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => EntRoleFromJSON(jsonValue));
    }

    /**
     * get role by ID
     * Get a role entity by ID
     */
    async getRole(requestParameters: GetRoleRequest): Promise<EntRole> {
        const response = await this.getRoleRaw(requestParameters);
        return await response.value();
    }

    /**
     * get user by ID
     * Get a user entity by ID
     */
    async getUserRaw(requestParameters: GetUserRequest): Promise<runtime.ApiResponse<EntUser>> {
        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling getUser.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/users/{id}`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => EntUserFromJSON(jsonValue));
    }

    /**
     * get user by ID
     * Get a user entity by ID
     */
    async getUser(requestParameters: GetUserRequest): Promise<EntUser> {
        const response = await this.getUserRaw(requestParameters);
        return await response.value();
    }

    /**
     * list booking entities
     * List booking entities
     */
    async listBookingRaw(requestParameters: ListBookingRequest): Promise<runtime.ApiResponse<Array<EntBooking>>> {
        const queryParameters: runtime.HTTPQuery = {};

        if (requestParameters.limit !== undefined) {
            queryParameters['limit'] = requestParameters.limit;
        }

        if (requestParameters.offset !== undefined) {
            queryParameters['offset'] = requestParameters.offset;
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/bookings`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => jsonValue.map(EntBookingFromJSON));
    }

    /**
     * list booking entities
     * List booking entities
     */
    async listBooking(requestParameters: ListBookingRequest): Promise<Array<EntBooking>> {
        const response = await this.listBookingRaw(requestParameters);
        return await response.value();
    }

    /**
     * list bookingtype entities
     * List bookingtype entities
     */
    async listBookingtypeRaw(requestParameters: ListBookingtypeRequest): Promise<runtime.ApiResponse<Array<EntBookingtype>>> {
        const queryParameters: runtime.HTTPQuery = {};

        if (requestParameters.limit !== undefined) {
            queryParameters['limit'] = requestParameters.limit;
        }

        if (requestParameters.offset !== undefined) {
            queryParameters['offset'] = requestParameters.offset;
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/bookingtypes`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => jsonValue.map(EntBookingtypeFromJSON));
    }

    /**
     * list bookingtype entities
     * List bookingtype entities
     */
    async listBookingtype(requestParameters: ListBookingtypeRequest): Promise<Array<EntBookingtype>> {
        const response = await this.listBookingtypeRaw(requestParameters);
        return await response.value();
    }

    /**
     * list cliententity entities
     * List cliententity entities
     */
    async listCliententityRaw(requestParameters: ListCliententityRequest): Promise<runtime.ApiResponse<Array<EntClientEntity>>> {
        const queryParameters: runtime.HTTPQuery = {};

        if (requestParameters.limit !== undefined) {
            queryParameters['limit'] = requestParameters.limit;
        }

        if (requestParameters.offset !== undefined) {
            queryParameters['offset'] = requestParameters.offset;
        }

        if (requestParameters.offset2 !== undefined) {
            queryParameters['offset'] = requestParameters.offset2;
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/cliententitys`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => jsonValue.map(EntClientEntityFromJSON));
    }

    /**
     * list cliententity entities
     * List cliententity entities
     */
    async listCliententity(requestParameters: ListCliententityRequest): Promise<Array<EntClientEntity>> {
        const response = await this.listCliententityRaw(requestParameters);
        return await response.value();
    }

    /**
     * list role entities
     * List role entities
     */
    async listRoleRaw(requestParameters: ListRoleRequest): Promise<runtime.ApiResponse<Array<EntRole>>> {
        const queryParameters: runtime.HTTPQuery = {};

        if (requestParameters.limit !== undefined) {
            queryParameters['limit'] = requestParameters.limit;
        }

        if (requestParameters.offset !== undefined) {
            queryParameters['offset'] = requestParameters.offset;
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/roles`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => jsonValue.map(EntRoleFromJSON));
    }

    /**
     * list role entities
     * List role entities
     */
    async listRole(requestParameters: ListRoleRequest): Promise<Array<EntRole>> {
        const response = await this.listRoleRaw(requestParameters);
        return await response.value();
    }

    /**
     * list user entities
     * List user entities
     */
    async listUserRaw(requestParameters: ListUserRequest): Promise<runtime.ApiResponse<Array<EntUser>>> {
        const queryParameters: runtime.HTTPQuery = {};

        if (requestParameters.limit !== undefined) {
            queryParameters['limit'] = requestParameters.limit;
        }

        if (requestParameters.offset !== undefined) {
            queryParameters['offset'] = requestParameters.offset;
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/users`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => jsonValue.map(EntUserFromJSON));
    }

    /**
     * list user entities
     * List user entities
     */
    async listUser(requestParameters: ListUserRequest): Promise<Array<EntUser>> {
        const response = await this.listUserRaw(requestParameters);
        return await response.value();
    }

    /**
     * update booking by ID
     * Update a booking entity by ID
     */
    async updateBookingRaw(requestParameters: UpdateBookingRequest): Promise<runtime.ApiResponse<EntBooking>> {
        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling updateBooking.');
        }

        if (requestParameters.booking === null || requestParameters.booking === undefined) {
            throw new runtime.RequiredError('booking','Required parameter requestParameters.booking was null or undefined when calling updateBooking.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/bookings/{id}`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'PUT',
            headers: headerParameters,
            query: queryParameters,
            body: EntBookingToJSON(requestParameters.booking),
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => EntBookingFromJSON(jsonValue));
    }

    /**
     * update booking by ID
     * Update a booking entity by ID
     */
    async updateBooking(requestParameters: UpdateBookingRequest): Promise<EntBooking> {
        const response = await this.updateBookingRaw(requestParameters);
        return await response.value();
    }

    /**
     * update bookingtype by ID
     * Update a bookingtype entity by ID
     */
    async updateBookingtypeRaw(requestParameters: UpdateBookingtypeRequest): Promise<runtime.ApiResponse<EntBookingtype>> {
        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling updateBookingtype.');
        }

        if (requestParameters.bookingtype === null || requestParameters.bookingtype === undefined) {
            throw new runtime.RequiredError('bookingtype','Required parameter requestParameters.bookingtype was null or undefined when calling updateBookingtype.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/bookingtypes/{id}`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'PUT',
            headers: headerParameters,
            query: queryParameters,
            body: EntBookingtypeToJSON(requestParameters.bookingtype),
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => EntBookingtypeFromJSON(jsonValue));
    }

    /**
     * update bookingtype by ID
     * Update a bookingtype entity by ID
     */
    async updateBookingtype(requestParameters: UpdateBookingtypeRequest): Promise<EntBookingtype> {
        const response = await this.updateBookingtypeRaw(requestParameters);
        return await response.value();
    }

    /**
     * update cliententity by ID
     * Update a cliententity entity by ID
     */
    async updateCliententityRaw(requestParameters: UpdateCliententityRequest): Promise<runtime.ApiResponse<EntClientEntity>> {
        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling updateCliententity.');
        }

        if (requestParameters.cliententity === null || requestParameters.cliententity === undefined) {
            throw new runtime.RequiredError('cliententity','Required parameter requestParameters.cliententity was null or undefined when calling updateCliententity.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/cliententitys/{id}`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'PUT',
            headers: headerParameters,
            query: queryParameters,
            body: EntClientEntityToJSON(requestParameters.cliententity),
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => EntClientEntityFromJSON(jsonValue));
    }

    /**
     * update cliententity by ID
     * Update a cliententity entity by ID
     */
    async updateCliententity(requestParameters: UpdateCliententityRequest): Promise<EntClientEntity> {
        const response = await this.updateCliententityRaw(requestParameters);
        return await response.value();
    }

    /**
     * update role by ID
     * Update a role entity by ID
     */
    async updateRoleRaw(requestParameters: UpdateRoleRequest): Promise<runtime.ApiResponse<EntRole>> {
        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling updateRole.');
        }

        if (requestParameters.role === null || requestParameters.role === undefined) {
            throw new runtime.RequiredError('role','Required parameter requestParameters.role was null or undefined when calling updateRole.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/roles/{id}`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'PUT',
            headers: headerParameters,
            query: queryParameters,
            body: EntRoleToJSON(requestParameters.role),
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => EntRoleFromJSON(jsonValue));
    }

    /**
     * update role by ID
     * Update a role entity by ID
     */
    async updateRole(requestParameters: UpdateRoleRequest): Promise<EntRole> {
        const response = await this.updateRoleRaw(requestParameters);
        return await response.value();
    }

    /**
     * update user by ID
     * Update a user entity by ID
     */
    async updateUserRaw(requestParameters: UpdateUserRequest): Promise<runtime.ApiResponse<EntUser>> {
        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling updateUser.');
        }

        if (requestParameters.user === null || requestParameters.user === undefined) {
            throw new runtime.RequiredError('user','Required parameter requestParameters.user was null or undefined when calling updateUser.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/users/{id}`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'PUT',
            headers: headerParameters,
            query: queryParameters,
            body: EntUserToJSON(requestParameters.user),
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => EntUserFromJSON(jsonValue));
    }

    /**
     * update user by ID
     * Update a user entity by ID
     */
    async updateUser(requestParameters: UpdateUserRequest): Promise<EntUser> {
        const response = await this.updateUserRaw(requestParameters);
        return await response.value();
    }

}