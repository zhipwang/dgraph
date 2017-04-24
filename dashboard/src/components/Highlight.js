import React from 'react';
import ReactDOM from 'react-dom';
import classnames from 'classnames';
import hljs from 'highlight.js/lib/highlight';

export default class Highlight extends React.Component {
  constructor(props) {
    super(props);

    this.highlightCode = this.highlightCode.bind(this);
  }

  componentDidMount() {
    this.highlightCode();
  }

  componentDidUpdate() {
    this.highlightCode();
  }

  highlightCode() {
    const domNode = ReactDOM.findDOMNode(this);
    hljs.highlightBlock(domNode);
  }

  render() {
    const { children, codeClass, preClass } = this.props;

    return (
      <pre className={classnames(preClass)}>
        <code className={classnames('json', codeClass)}>{children}</code></pre>
    );
  }
}
