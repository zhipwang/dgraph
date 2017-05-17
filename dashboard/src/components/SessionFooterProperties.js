import React from "react";
import classnames from "classnames";

import "../assets/css/SessionFooterProperties.css";

class SessionFooterProperties extends React.Component {
  constructor(props) {
    super(props);

    this.state = {
      isExpanded: false
    };
  }

  handleToggleExpand = e => {
    e.preventDefault();

    this.setState({
      isExpanded: !this.state.isExpanded
    });
  };

  render() {
    const { entity } = this.props;
    const { isExpanded } = this.state;

    const nodeProperties = JSON.parse(entity.title);

    // Nodes have facets and attrs keys.
    const isNode = Object.keys(nodeProperties).length !== 1;
    const attrs = nodeProperties.attrs || {};

    return (
      <div className="properties-container">
        <div className={classnames("properties", { expanded: isExpanded })}>
          <span className="label label-default entity-type-label">
            {isNode ? "Node" : "Edge"}
          </span>

          {isNode
            ? Object.keys(attrs).map(function(key, idx) {
                return (
                  <span className="property-pair" key={idx}>
                    <span className="property-key">
                      <span className="key-content">{key}</span>:
                    </span>
                    <span className="property-val">
                      {String(attrs[key])}
                    </span>
                  </span>
                );
              })
            : null}
        </div>

        <a
          href="#toggle-expand"
          onClick={this.handleToggleExpand}
          className="toggle-expand-btn"
        >
          {isExpanded
            ? <i className="fa fa-caret-up" />
            : <i className="fa fa-caret-down" />}
        </a>
      </div>
    );
  }
}

export default SessionFooterProperties;
