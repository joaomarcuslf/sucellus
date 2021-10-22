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
  const $loader = document.querySelector("#loader");
  $loader.classList.remove("is-hidden");

  const response = await fetch(url, {
    ...defaultHeaders,
    headers: {
      ...defaultHeaders.headers,
      ...headers,
    },
    method: 'GET',
  });

  setTimeout(() => {
    $loader.classList.add("is-hidden");
  }, 700)

  if (response.status >= 300) {
    throw await response.json();
  }

  return response.json();
};


const postJson = async (url, data, headers = {}) => {
  const $loader = document.querySelector("#loader");
  $loader.classList.remove("is-hidden");

  const response = await fetch(url, {
    ...defaultHeaders,
    headers: {
      ...defaultHeaders.headers,
      ...headers,
    },
    method: 'POST',
    body: JSON.stringify(data),
  });

  setTimeout(() => {
    $loader.classList.add("is-hidden");
  }, 700)

  if (response.status >= 300) {
    throw await response.json();
  }

  return response.json();
};

const deleteJson = async (url, headers = {}) => {
  const $loader = document.querySelector("#loader");
  $loader.classList.remove("is-hidden");

  const response = await fetch(url, {
    ...defaultHeaders,
    headers: {
      ...defaultHeaders.headers,
      ...headers,
    },
    method: 'DELETE',
  });

  setTimeout(() => {
    $loader.classList.add("is-hidden");
  }, 700)

  if (response.status >= 300) {
    throw await response.json();
  }

  return response.json();
};
