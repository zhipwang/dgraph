import React from "react";
import ReactDOM from 'react-dom';
import { connect } from "react-redux";
import screenfull from "screenfull";
import { Alert } from "react-bootstrap";

import NavbarContainer from "../containers/NavbarContainer";
import PreviousQueryListContainer from "./PreviousQueryListContainer";
import Editor from "./Editor";
import Response from "./Response";
import {
  updateFullscreen,
  getQuery,
  updateInitialQuery,
  queryFound,
  initialServerState
} from "../actions";

import "../assets/css/App.css";

class App extends React.Component {
  enterFullScreen = () => {
    if (!screenfull.enabled) {
      return;
    }

    const responseEl = ReactDOM.findDOMNode(this.refs.response);
    screenfull.request(responseEl);
  };

  render = () => {
    return (
      <div>
        <NavbarContainer />
        <div className="container-fluid">
          <div className="row justify-content-md-center">
            <div className="col-sm-12">
              <div className="col-sm-8 col-sm-offset-2">
                {!this.props.found &&
                  <Alert
                    ref={alert => {
                      this.alert = alert;
                    }}
                    bsStyle="danger"
                    onDismiss={() => {
                      this.props.queryFound(true);
                    }}
                  >
                    Couldn't find query with the given id.
                  </Alert>}
              </div>
              <div className="col-sm-5">
                <Editor />
                <PreviousQueryListContainer xs="hidden-xs" />
              </div>
              <div className="col-sm-7">
                <label style={{ marginLeft: "5px" }}> Response </label>
                {screenfull.enabled &&
                  <div
                    title="Enter full screen mode"
                    className="pull-right App-fullscreen"
                    onClick={this.enterFullScreen}
                  >
                    <span
                      className="App-fs-icon glyphicon glyphicon-glyphicon glyphicon-resize-full"
                    />
                  </div>}
                <Response ref="response" />
                <PreviousQueryListContainer xs="visible-xs-block" />
              </div>
            </div>
          </div>
        </div>
      </div>
    );
  };

  componentDidMount = () => {
    const { handleUpdateFullscreen } = this.props;
    this.props.initialServerState();
    let id = this.props.match.params.id;
    if (id !== undefined) {
      this.props.getQuery(id);
    }

    document.addEventListener(screenfull.raw.fullscreenchange, handleUpdateFullscreen);
  };

  componentWillUnmount = () => {
    const { handleUpdateFullscreen } = this.props;

    document.removeEventListener(screenfull.raw.fullscreenchange, handleUpdateFullscreen);
  }

  componentWillReceiveProps = nextProps => {
    if (!nextProps.found) {
      // Lets auto close the alert after 2 secs.
      setTimeout(
        () => {
          this.props.queryFound(true);
        },
        3000
      );
    }
  };
}

const mapStateToProps = state => ({
  found: state.share.found
});

const mapDispatchToProps = dispatch => ({
  handleUpdateFullscreen: () => {
    const fsState = screenfull.isFullscreen;
    dispatch(updateFullscreen(fsState));
  },
  getQuery: id => {
    dispatch(getQuery(id));
  },
  updateInitialQuery: () => {
    dispatch(updateInitialQuery());
  },
  queryFound: found => {
    dispatch(queryFound(found));
  },
  initialServerState: () => {
    dispatch(initialServerState());
  }
});

export default connect(mapStateToProps, mapDispatchToProps)(App);
