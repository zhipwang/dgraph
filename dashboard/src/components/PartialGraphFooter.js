import React from 'react';

const PartialGraphFooter = ({ partiallyRendered, onExpandNetwork, onCollapseNetwork }) => {
  return (
    <div className="footer">
      {partiallyRendered ?
        <div>
          Only a subset of the graph was loaded. Double click on a leaf node to expand its child nodes, or <a href="#expand" onClick={onExpandNetwork}>exand next 500 nodes</a>
        </div>
        :
        <div><a href="#collapse" onClick={onCollapseNetwork}>Collapse nodes</a></div>
      }
    </div>
  );
};
export default PartialGraphFooter;
