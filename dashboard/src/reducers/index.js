import { combineReducers } from "redux";

import scratchpad from "./scratchpad";
import frames from "./frames";

const rootReducer = combineReducers({
    scratchpad,
    frames
});

export default rootReducer;
