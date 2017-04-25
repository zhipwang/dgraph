import React from 'react';
import classnames from 'classnames';

import SessionQueryTab from './SessionQueryTab';
import SessionGraphTab from './SessionGraphTab';
import SessionJSONTab from './SessionJSONTab';
import SessionTreeTab from './SessionTreeTab';
import GraphIcon from './GraphIcon';
import TreeIcon from './TreeIcon';

class SessionItem extends React.Component {
  constructor(props) {
    super(props);

    this.state = {
      isShown: false,
      // tabs: query, graph, tree, json
      currentTab: 'graph'
    };
  }

  componentDidMount() {
    setTimeout(() => {
      this.setState({ isShown: true });
    }, 50);
  }

  navigateTab = (tabName) => {
    this.setState({
      currentTab: tabName
    });
  }

  render() {
    const { session } = this.props;
    const { isShown, currentTab } = this.state;

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
            <SessionQueryTab session={session} active={currentTab === 'query'} />
            <SessionGraphTab session={session} active={currentTab === 'graph'} />
            <SessionTreeTab session={session} active={currentTab === 'tree'} />
            <SessionJSONTab session={session} active={currentTab === 'json'} />
          </div>
        </div>
      </li>
    );
  }
}

export default SessionItem;
