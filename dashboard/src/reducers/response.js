import { UPDATE_RESPONSE } from "../actions/response";

const defaultState = {
	items: []
};

const frames = (state = defaultState, action) => {
	switch (action.type) {
		case UPDATE_RESPONSE:
			console.log(action);
			return {
				...state,
				items: [
					{
						id: action.id,
						response: action.response
					},
					...state.items
				]
			};
		default:
			return state;
	}
};

export default frames;
