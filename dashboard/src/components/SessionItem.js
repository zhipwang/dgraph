import React from 'react';
import classnames from 'classnames';

import SessionQueryTab from './SessionQueryTab';
import SessionGraphTab from './SessionGraphTab';
import SessionJSONTab from './SessionJSONTab';
import SessionTreeTab from './SessionTreeTab';
import SessionFooter from './SessionFooter';
import GraphIcon from './GraphIcon';
import TreeIcon from './TreeIcon';


class SessionItem extends React.Component {
  constructor(props) {
    super(props);

    this.state = {
      // tabs: query, graph, tree, json
      currentTab: 'graph',
      graphRenderStart: null,
      graphRenderEnd: null,
      treeRenderStart: null,
      treeRenderEnd: null,
      currentNode: null
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
    this.setState({ currentNode: node });
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
    const { isShown, currentTab, currentNode } = this.state;

    return (
      <li className={classnames('session-item', { shown: isShown })}>
        <div className="header">
          <div className="actions">
            <a href="#discard" className="action">
              <i className="fa fa-close" />
            </a>
            <a href="#fullscreen" className="action">
              <i className="fa fa-expand" />
            </a>
            <a href="#share" className="action">
              <i className="fa fa-share-alt" />
            </a>
          </div>
        </div>

        <div className="body">
          <div className="sidebar">
            <ul className="sidebar-nav">
              <li>
                <a
                  href="#query"
                  className={classnames('sidebar-nav-item', { active: currentTab === 'query' })}
                  onClick={this.navigateTab.bind(this, 'query')}
                >
                  <i className="icon fa fa-search" />
                  <span className="menu-label">Query</span>

                </a>
              </li>
              <li>
                <a
                  href="#graph"
                  className={classnames('sidebar-nav-item', { active: currentTab === 'graph' })}
                  onClick={this.navigateTab.bind(this, 'graph')}
                >
                  <GraphIcon />
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
                  <TreeIcon />
                  <span className="menu-label">Tree</span>

                </a>
              </li>
              <li>
                <a
                  href="#tree"
                  className={classnames('sidebar-nav-item', { active: currentTab === 'json' })}
                  onClick={this.navigateTab.bind(this, 'json')}
                >
                  <i className="icon fa fa-code" />
                  <span className="menu-label">JSON</span>

                </a>
              </li>
            </ul>
          </div>

          <div className="main">
            <SessionQueryTab
              session={session}
              active={currentTab === 'query'}
             />
            <SessionGraphTab
              session={session}
              active={currentTab === 'graph'}
              onBeforeGraphRender={this.handleBeforeGraphRender}
              onGraphRendered={this.handleGraphRendered}
              onNodeSelected={this.handleNodeSelected}
              currentNode={currentNode}
            />
            <SessionTreeTab
              session={session}
              active={currentTab === 'tree'}
              onBeforeTreeRender={this.handleBeforeTreeRender}
              onTreeRendered={this.handleTreeRendered}
              onNodeSelected={this.handleNodeSelected}
              currentNode={currentNode}
            />
            <SessionJSONTab
              session={session}
              active={currentTab === 'json'}
            />

            <SessionFooter
              session={session}
              currentTab={currentTab}
              currentNode={currentNode}
              graphRenderTime={this.getGraphRenderTime()}
              treeRenderTime={this.getTreeRenderTime()}
            />
          </div>
        </div>
      </li>
    );
  }
}

export default SessionItem;
