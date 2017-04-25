import React, { Component } from "react";
import { connect } from "react-redux";
import _ from "lodash/object";

import { renderNetwork } from '../lib/graph';

import "../assets/css/Graph.css";
import "vis/dist/vis.min.css";


class GraphContainer extends Component {
    constructor(props: Props) {
        super(props);

        this.state = {
            selectedNode: false,
        };
    }

    componentDidMount() {
      const { response, treeView } = this.props;

      renderNetwork({
        nodes: response.nodes,
        edges: response.edges,
        allNodes: response.allNodes,
        allEdges: response.allEdges,
        containerEl: this.refs.graph,
        treeView,
      });
    }

    render() {
        return (
            <div ref="graph" className="content" />
        );
    }
}

const mapStateToProps = state => ({
});

export default connect(mapStateToProps, null, null, { withRef: true })(
    GraphContainer
);
