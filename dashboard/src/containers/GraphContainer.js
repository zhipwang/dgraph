import React, { Component } from "react";
import vis from 'vis';
import { connect } from "react-redux";
import _ from "lodash/object";

import { renderNetwork } from '../lib/graph';
import Label from '../components/Label';

import "../assets/css/Graph.css";
import "vis/dist/vis.min.css";

const doubleClickTime = 0;
const threshold = 200;

class GraphContainer extends Component {
    constructor(props: Props) {
        super(props);

        this.state = {
          renderTime: null
        };
    }

    componentDidMount() {
      const { response, treeView, onRendered } = this.props;

      const { renderTime, network } = renderNetwork({
        nodes: response.nodes,
        edges: response.edges,
        allNodes: response.allNodes,
        allEdges: response.allEdges,
        containerEl: this.refs.graph,
        treeView,
      });

      this.configNetworkBehavior(network);

      if (onRendered) {
        onRendered(renderTime);
      }
    }

    /**
     * configNetworkBehavior configures the custom behaviors for a vis.Network
     */
    configNetworkBehavior = (network) => {
      const { response: { nodes, edges }, onNodeSelected } = this.props;
      const data = {
        nodes: new vis.DataSet(nodes),
        edges: new vis.DataSet(edges)
    };

      network.on("click", (params, allNode) => {
        const t0 = new Date();

        if (t0 - doubleClickTime > threshold) {
          setTimeout(
           () => {
              if (t0 - doubleClickTime < threshold) {
                return;
              }

              if (params.nodes.length > 0) {
                const nodeUid = params.nodes[0];
                const currentNode = data.nodes.get(nodeUid);

                onNodeSelected(currentNode);
              } else if (params.edges.length > 0) {
                const edgeUid = params.edges[0];
                const currentEdge = data.edges.get(edgeUid);

                onNodeSelected(currentEdge);
              } else {
                onNodeSelected(null);
              }
            },
            threshold
          );
        }
      });
    }

    render() {
        const { response, currentNode } = this.props;

        return (
          <div className="graph-container content">
            <div className="labels">
              {response.plotAxis.map((label, i) => {
                return (
                  <Label
                    key={i}
                    color={label.color}
                    pred={label.pred}
                    label={label.label}
                  />
                );
              })}
            </div>
            <div ref="graph" className="graph" />
            <div>

            </div>
          </div>
        );
    }
}

const mapStateToProps = state => ({
});

export default connect(mapStateToProps, null, null, { withRef: true })(
    GraphContainer
);
