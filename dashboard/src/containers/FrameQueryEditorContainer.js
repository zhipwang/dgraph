import { connect } from "react-redux";

import FrameQueryEditor from "../components/FrameQueryEditor";
import {
  runQuery
} from "../actions";

const mapStateToProps = state => ({
});

const mapDispatchToProps = dispatch => ({
  handleRunQuery(query, frameId) {
    dispatch(runQuery(query, frameId));
  }
});

export default connect(mapStateToProps, mapDispatchToProps)(FrameQueryEditor);
