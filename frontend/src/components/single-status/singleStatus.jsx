import { useState } from 'react';
import axios from 'axios';
import './singleStatus.css';

function StatusComponent({ url, name }) {
  const [status, setStatus] = useState(null);

  const fetchStatus = () => {
    axios.get(url)
      .then(response => {
        setStatus(response.data);
      })
      .catch(error => {
        console.error(`Error fetching ${name} status:`, error);
      });
  };

  return (
    <div className="status-container">
      <button onClick={fetchStatus}>Fetch {name} Status</button>
      <div>
        <p>Status Code: {status ? status.statusCode : "-"}</p>
        <p>Request Duration: {status ? status.duration : "0"} ms</p>
        <p>Date: {status ? status.date : "-"}</p>
      </div>
    </div>
  );
}

export default StatusComponent;
