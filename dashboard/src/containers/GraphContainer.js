import React, { Component } from "react";
import vis from 'vis';
import { connect } from "react-redux";
import _ from "lodash/object";

import { renderNetwork } from '../lib/graph';
import Label from '../components/Label';
import { outgoingEdges } from './Helpers';

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
      const {
        response: { allNodes, allEdges },
        onNodeSelected
      } = this.props;
      const { data } = network.body;
      const allEdgeSet = new vis.DataSet(allEdges);
      const allNodeSet = new vis.DataSet(allNodes);

      function multiLevelExpand(nodeId) {
        let nodes = [nodeId], nodeStack = [nodeId], adjEdges = [];
        while (nodeStack.length !== 0) {
            let nodeId = nodeStack.pop();

            let outgoing = outgoingEdges(nodeId, allEdgeSet),
                adjNodeIds = outgoing.map(function(edge) {
                    return edge.to;
                });

            nodeStack = nodeStack.concat(adjNodeIds);
            nodes = nodes.concat(adjNodeIds);
            adjEdges = adjEdges.concat(outgoing);
            if (adjNodeIds.length > 3) {
                break;
            }
        }
        data.nodes.update(allNodeSet.get(nodes));
        data.edges.update(adjEdges);
      }

      network.on("click", (params) => {
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

      network.on("doubleClick", (params) => {
        if (params.nodes && params.nodes.length > 0) {
            const clickedNodeUid = params.nodes[0];
            const clickedNode = data.nodes.get(clickedNodeUid);

            network.unselectAll();
            onNodeSelected(clickedNode);

            const outgoing = outgoingEdges(clickedNodeUid, data.edges);
            const allOutgoingEdges = outgoingEdges(clickedNodeUid, allEdgeSet);
            const expanded = outgoing.length > 0 || allOutgoingEdges.length === 0;

            let adjacentNodeIds: Array<string> = allOutgoingEdges.map(
                function(edge) {
                    return edge.to;
                }
            );

            let adjacentNodes = allNodeSet.get(adjacentNodeIds);

            // TODO -See if we can set a meta property to a node to know that its
            // expanded or closed and avoid this computation.
            if (expanded) {
                // Collapse all child nodes recursively.
                let allEdges = outgoing.map(function(edge) {
                    return edge.id;
                });

                let allNodes = adjacentNodes.slice();

                while (adjacentNodeIds.length > 0) {
                    let node = adjacentNodeIds.pop();
                    let connectedEdges = outgoingEdges(node, data.edges);

                    let connectedNodes = connectedEdges.map(function(edge) {
                        return edge.to;
                    });

                    allNodes = allNodes.concat(connectedNodes);
                    allEdges = allEdges.concat(connectedEdges);
                    adjacentNodeIds = adjacentNodeIds.concat(connectedNodes);
                }

                data.nodes.remove(allNodes);
                data.edges.remove(allEdges);
            } else {
                multiLevelExpand(clickedNodeUid);
                if (data.nodes.length === allNodeSet.length) {
                    // TODO: what is partial?
                    // dispatch(updatePartial(false));
                }
            }
        }
      });
    }

    render() {
        const { response } = this.props;

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
          </div>
        );
    }
}

const mapStateToProps = state => ({
});

export default connect(mapStateToProps, null, null, { withRef: true })(
    GraphContainer
);
