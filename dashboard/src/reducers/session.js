import { RECEIVE_SESSION } from '../actions/session';

const defaultState = {
  isLoading: false,
  items: []
}

const session = (state = defaultState, action) => {
    switch (action.type) {
        case RECEIVE_SESSION:
            return {
                ...state,
                items: [ action.session, ...state.items ]
            };
        default:
            return state;
    }
};

export default session;
