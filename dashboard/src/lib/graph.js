/**
 * graph helpers
 */

import vis from "vis";
import _ from "lodash";

/**
 * renderNetwork renders a vis.Network within the containerEl
 * containerEl {HTMLElement}
 */
export function renderNetwork({ nodes, edges, treeView, allNodes, allEdges, containerEl}) {
    var data = {
        nodes: new vis.DataSet(nodes),
        edges: new vis.DataSet(edges)
    };
    var options = {
        nodes: {
            shape: "circle",
            scaling: {
                max: 20,
                min: 20,
                label: {
                    enabled: true,
                    min: 14,
                    max: 14
                }
            },
            font: {
                size: 16
            },
            margin: {
                top: 25
            }
        },
        height: "100%",
        width: "100%",
        interaction: {
            hover: true,
            keyboard: {
                enabled: true,
                bindToWindow: false
            },
            navigationButtons: true,
            tooltipDelay: 1000000,
            hideEdgesOnDrag: true,
            zoomView: false
        },
        layout: {
            randomSeed: 42,
            improvedLayout: false
        },
        physics: {
            stabilization: {
                fit: true,
                updateInterval: 5,
                iterations: 20
            },
            barnesHut: {
                damping: 0.7
            }
        }
    };

    if (data.nodes.length < 100) {
        _.merge(options, {
            physics: {
                stabilization: {
                    iterations: 200,
                    updateInterval: 50
                }
            }
        });
    }

    if (treeView) {
        Object.assign(options, {
            layout: {
                hierarchical: {
                    sortMethod: "directed"
                }
            },
            physics: {
                // Otherwise there is jittery movement (existing nodes move
                // horizontally which doesn't look good) when you expand some nodes.
                enabled: false,
                barnesHut: {}
            }
        });
    }

    const network = new vis.Network(containerEl, data, options);
    return network;
}
