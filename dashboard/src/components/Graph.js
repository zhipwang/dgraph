import React from "react";
// import { Nav, NavItem } from "react-bootstrap";

// import Label from "../components/Label";

import "../assets/css/Graph.css";

class Graph extends React.Component {
  render() {
    let {
        text,
        success,
        // plotAxis,
        fs,
        isFetching,
        // json,
        // selectTab,
        selectedTab,
        // saveGraphRef
    } = this.props,
        graphClass = fs ? "Graph-fs" : "Graph-s",
        bgColor,
        hourglass = isFetching ? "Graph-hourglass" : "",
        graphHeight = fs ? "Graph-full-height" : "Graph-fixed-height",
        showGraph = selectedTab === "1" ? "" : "Graph-hide";
        // showJSON = selectedTab === "2" ? "" : "Graph-hide";

    if (success) {
        if (text !== "") {
            bgColor = "Graph-success";
        } else {
            bgColor = "";
        }
    } else if (text !== "") {
        bgColor = "Graph-error";
    }

    return (
        <div className="Graph-wrapper">
            <div className="Graph-outer">
                <div id="graph">
                  
                </div>
            </div>
        </div>
    );
  }
}

export default Graph;
