export const UPDATE_RESPONSE = "response/UPDATE_RESPONSE";

export function updateResponse({ id, response }) {
	return {
		type: UPDATE_RESPONSE,
		id,
		response
	};
}
