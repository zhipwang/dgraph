import React from 'react';
import classnames from 'classnames';

import FrameQueryTab from './FrameQueryTab';
import FrameMessageTab from './FrameMessageTab';

class FrameSuccess extends React.Component {
  constructor(props) {
    super(props);

    this.state = {
      currentTab: 'message'
    };
  }

  navigateTab = (tabName, e) => {
    e.preventDefault();

    this.setState({
      currentTab: tabName
    });
  }

  render() {
    const { data } = this.props;
    const { currentTab } = this.state;

    return (
      <div className="body">
        <div className="content">
          <div className="sidebar">
            <ul className="sidebar-nav">
              <li>
                <a
                  href="#query"
                  className={classnames('sidebar-nav-item', { active: currentTab === 'message' })}
                  onClick={this.navigateTab.bind(this, 'message')}
                >
                  <div className="icon-container">
                    <i className="icon fa fa-check-circle" />
                  </div>
                  <span className="menu-label">Message</span>
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
            {currentTab === 'message' ?
              <FrameMessageTab message={data.message} /> :null}
            {currentTab === 'response' ?
              <FrameQueryTab query={data.response} /> :null}
          </div>
        </div>
      </div>
    );
  }
}

export default FrameSuccess;
