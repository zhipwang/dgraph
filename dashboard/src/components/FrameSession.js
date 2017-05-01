import React from 'react';
import classnames from 'classnames';

import SessionQueryTab from './SessionQueryTab';
import SessionGraphTab from './SessionGraphTab';
import SessionJSONTab from './SessionJSONTab';
import SessionTreeTab from './SessionTreeTab';
import SessionFooter from './SessionFooter';
import GraphIcon from './GraphIcon';
import TreeIcon from './TreeIcon';


class FrameSession extends React.Component {
  constructor(props) {
    super(props);

    this.state = {
      // tabs: query, graph, tree, json
      currentTab: 'graph',
      graphRenderStart: null,
      graphRenderEnd: null,
      treeRenderStart: null,
      treeRenderEnd: null,
      selectedNode: null,
      hoveredNode: null,
      isTreePartial: false
    };
  }

  handleBeforeGraphRender = () => {
    this.setState({ graphRenderStart: new Date() });
  }

  handleGraphRendered = () => {
    this.setState({ graphRenderEnd: new Date() });
  }

  handleBeforeTreeRender = () => {
    this.setState({ treeRenderStart: new Date() });
  }

  handleTreeRendered = () => {
    this.setState({ treeRenderEnd: new Date() });
  }

  handleNodeSelected = (node) => {
    const { selectedNode } = this.state;

    if (!node) {
      this.setState({
        selectedNode: null,
        hoveredNode: null
      });
      return;
    }

    this.setState({ selectedNode: node });
  }

  handleNodeHovered = (node) => {
    this.setState({ hoveredNode: node });
  }

  navigateTab = (tabName, e) => {
    e.preventDefault();

    this.setState({
      currentTab: tabName
    });
  }

  getGraphRenderTime = () => {
    const { graphRenderStart, graphRenderEnd } = this.state;
    if (!graphRenderStart || !graphRenderEnd) {
      return
    }

    return graphRenderEnd.getTime() - graphRenderStart.getTime();
  }

  getTreeRenderTime = () => {
    const { treeRenderStart, treeRenderEnd } = this.state;
    if (!treeRenderStart || !treeRenderEnd) {
      return
    }

    return treeRenderEnd.getTime() - treeRenderStart.getTime();
  }

  render() {
    const { session } = this.props;
    const { currentTab, selectedNode, hoveredNode } = this.state;

    return (
      <div className="body">
        <div className="content">
          <div className="sidebar">
            <ul className="sidebar-nav">
              <li>
                <a
                  href="#query"
                  className={classnames('sidebar-nav-item', { active: currentTab === 'query' })}
                  onClick={this.navigateTab.bind(this, 'query')}
                >
                  <div className="icon-container">
                    <i className="icon fa fa-search" />
                  </div>
                  <span className="menu-label">Query</span>

                </a>
              </li>
              <li>
                <a
                  href="#graph"
                  className={classnames('sidebar-nav-item', { active: currentTab === 'graph' })}
                  onClick={this.navigateTab.bind(this, 'graph')}
                >
                  <div className="icon-container">
                    <GraphIcon />
                  </div>
                  <span className="menu-label">
                    Graph
                  </span>
                </a>
              </li>
              <li>
                <a
                  href="#tree"
                  className={classnames('sidebar-nav-item', { active: currentTab === 'tree' })}
                  onClick={this.navigateTab.bind(this, 'tree')}
                >
                  <div className="icon-container">
                    <TreeIcon />
                  </div>
                  <span className="menu-label">Tree</span>

                </a>
              </li>
              <li>
                <a
                  href="#tree"
                  className={classnames('sidebar-nav-item', { active: currentTab === 'json' })}
                  onClick={this.navigateTab.bind(this, 'json')}
                >
                  <div className="icon-container">
                    <i className="icon fa fa-code" />
                  </div>

                  <span className="menu-label">JSON</span>

                </a>
              </li>
            </ul>
          </div>

          <div className="main">
            {currentTab === 'query' ?
              <SessionQueryTab
                session={session}
               /> :null}

             {currentTab === 'graph' ?
               <SessionGraphTab
                 session={session}
                 onBeforeGraphRender={this.handleBeforeGraphRender}
                 onGraphRendered={this.handleGraphRendered}
                 onNodeSelected={this.handleNodeSelected}
                 onNodeHovered={this.handleNodeHovered}
                 selectedNode={selectedNode}
                 hoveredNode={hoveredNode}
               /> : null}

             {currentTab === 'tree' ?
               <SessionTreeTab
                 session={session}
                 onBeforeTreeRender={this.handleBeforeTreeRender}
                 onTreeRendered={this.handleTreeRendered}
                 onNodeSelected={this.handleNodeSelected}
                 onNodeHovered={this.handleNodeHovered}
                 selectedNode={selectedNode}
               /> : null}

             {currentTab === 'json' ?
               <SessionJSONTab
                 session={session}
               /> : null}
          </div>
        </div>

        <SessionFooter
          session={session}
          currentTab={currentTab}
          selectedNode={selectedNode}
          hoveredNode={hoveredNode}
          graphRenderTime={this.getGraphRenderTime()}
          treeRenderTime={this.getTreeRenderTime()}
        />
      </div>
    );
  }
}

export default FrameSession;
