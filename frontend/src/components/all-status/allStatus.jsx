import React, { useState, useEffect } from 'react'
import axios from 'axios'
import './allStatus.css'

function Status() {
  const [status, setStatus] = useState({ loading: true, data: null, error: null })

  useEffect(() => {
    const fetchData = () => {
        axios.get(`/api/v1/all-status?timestamp=${new Date().getTime()}`)
        .then(response => {
          setStatus({ loading: false, data: response.data, error: null })
        })
        .catch(error => {
          setStatus({ loading: false, data: null, error: error.message })
        })
    }
    fetchData();

    // Fetches status data every minute
    const interval = setInterval(() => {
      fetchData();
    }, 60000); 

    // Clears the interval on component unmount
    return () => clearInterval(interval); 
  }, [])

  // helper funtion to format date for readability
  function formatDate(dateString) {
    const options = { year: 'numeric', month: 'long', day: 'numeric', hour: '2-digit', minute: '2-digit', hour12: true };
    return new Intl.DateTimeFormat('en-US', options).format(new Date(dateString));
  }
  
  if (status.loading) {
    return <span>Loading...</span>
  } else if (status.error) {
    return <span>Error: {status.error}</span>
  }

  return (
    <div className="all-status-container">
      <h2>All Statuses</h2>
      {status.data.map((siteStatus, index) => (
        <div key={index} className="status-section">
          <h3>{siteStatus.url}</h3>
          <div className="status-elements">
            <p>Status Code: {siteStatus.statusCode}</p>
            <p>Request Duration: {siteStatus.duration} ms</p>
            <p>Date: {formatDate(siteStatus.date)}</p>
          </div>
        </div>
      ))}
    </div>
  );
}

export default Status
