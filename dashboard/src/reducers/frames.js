import { RECEIVE_FRAME } from '../actions/frames';

const defaultState = {
  items: []
}

const frames = (state = defaultState, action) => {
    switch (action.type) {
        case RECEIVE_FRAME:
            return {
                ...state,
                items: [ action.frame, ...state.items ]
            };
        default:
            return state;
    }
};

export default frames;
