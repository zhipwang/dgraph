import { combineReducers } from "redux";
import previousQueries from "./previousQueries";
import query from "./query";
import response from "./response";
import interaction from "./interaction";
import latency from "./latency";
import scratchpad from "./scratchpad";
import share from "./share";
import session from "./session";

const rootReducer = combineReducers({
    query,
    previousQueries,
    response,
    interaction,
    latency,
    scratchpad,
    share,
    session
});

export default rootReducer;
