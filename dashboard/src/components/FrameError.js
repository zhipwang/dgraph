import React from 'react';
import classnames from 'classnames';

import FrameQueryTab from './FrameQueryTab';
import FrameErrorTab from './FrameErrorTab';

class FrameError extends React.Component {
  constructor(props) {
    super(props);

    this.state = {
      // tabs: 'error', 'response', 'query'
      currentTab: 'error'
    }
  }

  navigateTab = (tabName, e) => {
    e.preventDefault();

    this.setState({
      currentTab: tabName
    });
  }

  render() {
    const { data: { message, query, response } } = this.props;
    const { currentTab } = this.state;

    return (
      <div className="body">
        <div className="content">
          <div className="sidebar">
            <ul className="sidebar-nav">
              <li>
                <a
                  href="#query"
                  className={classnames('sidebar-nav-item', { active: currentTab === 'error' })}
                  onClick={this.navigateTab.bind(this, 'error')}
                >
                  <div className="icon-container">
                    <i className="icon fa fa-warning" />
                  </div>
                  <span className="menu-label">Error</span>

                </a>
              </li>
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
                  href="#tree"
                  className={classnames('sidebar-nav-item', { active: currentTab === 'response' })}
                  onClick={this.navigateTab.bind(this, 'response')}
                >
                  <div className="icon-container">
                    <i className="icon fa fa-code" />
                  </div>

                  <span className="menu-label">Response</span>

                </a>
              </li>
            </ul>
          </div>

          <div className="main">
            {currentTab === 'error' ?
              <FrameErrorTab message={message} /> :null}
            {currentTab === 'query' ?
              <FrameQueryTab query={query} /> :null}
            {currentTab === 'response' ?
              <FrameQueryTab query={response} /> :null}
          </div>
        </div>

        <div className="footer error-footer">
          <i className="fa fa-warning error-mark" /> <span className="result-message">Error occurred</span>
        </div>
      </div>
    );
  }
}

export default FrameError;
