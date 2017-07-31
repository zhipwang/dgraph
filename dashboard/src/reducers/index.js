import { combineReducers } from "redux";

import frames from "./frames";
import connection from "./connection";
import response from "./response";

const rootReducer = combineReducers({
	frames,
	connection,
	response
});

export default rootReducer;
