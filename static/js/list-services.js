var template = ({ name, url, port, language, pooling_interval, status }) => {
  return `
  <tr>
    <td>${name}</td>
    <td><a href="${url}" target="blank">${url}</a></td>
    <td>${port}</td>
    <td>${language}</td>
    <td>${pooling_interval}</td>
    <td>${status}</td>
  </tr>
  `
}

var renderServices = async (print = true) => {
  const $target = document.querySelector('#servicesListTable')

  try {
    const response  = await fetchJson("/api/services")

    const services = response.data

    if (services.length > 0) {
      let html = []

      html = services.map(service => {
        return template(service)
      })

      $target.innerHTML = html.join('')
      print && notyf.success("Loaded services!");
    } else {
      $target.innerHTML = '<tr><td colspan="6">No services found</td></tr>'
      print && notyf.warn("No services to be displayed!");
    }
  } catch(ex) {
    console.error(ex);

    $target.innerHTML = '<tr><td colspan="6">No services found</td></tr>'
    print && notyf.error("Error loading services!");
  }
}


document.addEventListener('DOMContentLoaded', async () => {
  const $poolingInput = document.querySelector('#poolingInput');
  const $addServiceBtn = document.querySelector('#addServiceBtn');
  const $startServicesBtn = document.querySelector('#startServicesBtn');
  const $stopServicesBtn = document.querySelector('#stopServicesBtn');

  $poolingInput.checked = true;

  await renderServices();

  setInterval(async () => {
    if ($poolingInput.checked) {
      await renderServices();
    }
  }, 10 * 1000);

  $addServiceBtn.addEventListener('click', async () => {
    window.location.href = '/services/new'
  });

  $startServicesBtn.addEventListener('click', async () => {
    try {
      await postJson('/api/services/start')

      notyf.success("Starting all services!");

      await renderServices(false);
    } catch(ex) {
      console.error(ex);
      notyf.error("Error starting services!");
    }
  });

  $stopServicesBtn.addEventListener('click', async () => {
    try {
      await postJson('/api/services/stop')

      notyf.success("Stopping all services!");

      await renderServices(false);
    } catch(ex) {
      console.error(ex);
      notyf.error("Error stopping services!");
    }
  });
});
