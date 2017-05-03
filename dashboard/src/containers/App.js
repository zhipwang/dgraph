import React from "react";
import { connect } from "react-redux";
import screenfull from "screenfull"

import Sidebar from '../components/Sidebar';
import EditorPanel from '../components/EditorPanel';
import FrameList from '../components/FrameList';
import {
  runQuery
} from "../actions";
import { discardFrame } from '../actions/frames';
import { readCookie, eraseCookie } from './Helpers';

import "../assets/css/App.css";

class App extends React.Component {
  componentDidMount = () => {
    const { handleSelectQuery, handleRunQuery, handleUpdateFullscreen } = this.props;

    let id = this.props.match.params.id;
    if (id !== undefined) {
      this.props.getQuery(id);
    }

    // If playQuery cookie is set, run the query and erase the cookie
    // The cookie is used to communicate the query string between docs and play
    const playQuery = readCookie('playQuery');
    if (playQuery) {
      const queryString = decodeURI(playQuery);
      handleSelectQuery(queryString);
      handleRunQuery(queryString).then(() => {
        eraseCookie('playQuery', { crossDomain: true });
      });
    }

    document.addEventListener(screenfull.raw.fullscreenchange, handleUpdateFullscreen);
  };

  componentWillUnmount = () => {
    const { handleUpdateFullscreen } = this.props;

    document.removeEventListener(screenfull.raw.fullscreenchange, handleUpdateFullscreen);
  }

  toggleSidebarOpen = () => {
    this.setState({ sidebarOpen: !this.state.sidebarOpen });
  }

  render = () => {
    const { handleRunQuery, handleDiscardFrame, frames } = this.props;

    return (
      <div className="app-layout">
        <Sidebar />
        <div className="main-content">
          <div className="container-fluid">
            <div className="row justify-content-md-center">
              <div className="col-sm-12">
                <EditorPanel
                  onRunQuery={handleRunQuery}
                />
              </div>

              <div className="col-sm-12">
                <FrameList
                  frames={frames}
                  onDiscardFrame={handleDiscardFrame}
                />
              </div>
            </div>
          </div>
        </div>
      </div>
    );
  };
}

const mapStateToProps = state => ({
  frames: state.frames.items
});

const mapDispatchToProps = dispatch => ({
  handleRunQuery: (query, done = () => {}) => {
    return dispatch(runQuery(query))
      .then(done);
  },
  handleDiscardFrame(frameID) {
    dispatch(discardFrame(frameID));
  }
});

export default connect(mapStateToProps, mapDispatchToProps)(App);
