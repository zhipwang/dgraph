import { RECEIVE_FRAME, DISCARD_FRAME } from '../actions/frames';

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
    case DISCARD_FRAME:
      return {
        ...state,
        items: state.items.filter(item => item.id !== action.frameID)
      }
    default:
      return state;
  }
};

export default frames;
