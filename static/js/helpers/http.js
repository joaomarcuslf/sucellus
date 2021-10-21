const defaultHeaders = {
  cache: 'no-cache',
  credentials: 'same-origin',
  method: 'GET',
  mode: 'cors',
  redirect: 'follow',
  referrer: 'no-referrer',
  headers: {
    'Content-Type': 'application/json',
  },
};

const fetchJson = async (url, headers = {}) => {
  const response = await fetch(url, {
    ...defaultHeaders,
    headers: {
      ...defaultHeaders.headers,
      ...headers,
    },
    method: 'GET',
  });

  return response.json();
};


const postJson = async (url, data, headers = {}) => {
  const response = await fetch(url, {
    ...defaultHeaders,
    headers: {
      ...defaultHeaders.headers,
      ...headers,
    },
    method: 'POST',
    body: JSON.stringify(data),
  });

  return response.json();
};
