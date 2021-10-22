const getInput = (container) => {
  return container.querySelector(".control select") || container.querySelector(".control textarea") || container.querySelector(".control .control > *") || container.querySelector(".control > *")
}
document.addEventListener('DOMContentLoaded', async () => {
  const $createServiceForm = document.querySelector('#createServiceForm');

  $createServiceForm.addEventListener("submit", async (evt) => {
    evt.preventDefault();

    const fields = {
      name: $createServiceForm.querySelector("#name"),
      port: $createServiceForm.querySelector("#port"),
      poolinginterval: $createServiceForm.querySelector("#poolinginterval"),
      language: $createServiceForm.querySelector("#language"),
      url: $createServiceForm.querySelector("#url"),
      envvars: $createServiceForm.querySelector("#envvars")
    }

    const envVars = getInput(fields.envvars).value.trim().split("\n").reduce(function (acc, cur) {
        const [key, value] = cur.split("=")
        acc[key] = value.replace(/;/g, '')
        return acc
      },
      {}
    );


    const data = {
      name: (getInput(fields.name)).value,
      port: parseInt(getInput(fields.port).value),
      pooling_interval: parseInt(getInput(fields.poolinginterval).value),
      language: (getInput(fields.language)).value,
      url: (getInput(fields.url)).value,
      env_vars: envVars,
    };

    try {
      Object.keys(fields).forEach(key => {
        const $container = fields[key];

        const $input = $container.querySelector(".control .control > *") || $container.querySelector(".control > *");
        const $help = $container.querySelector(".help");

        $container.classList.remove("has-error");
        $input.classList.remove("is-danger");

        $help.innerHTML = "";
      });

      await postJson('/api/services', data);

      window.location.href = '/services';

      notyf.success("Starting all services!");
    } catch(ex) {

      if (ex.message && ex.message.includes("FieldName")) {
        let fieldName = ex.message.split("\n")[1].split("FieldName: ")[1];
        let fieldMessage = ex.message.split("\n")[3].split("FieldValidation: ")[1];

        fieldMessage = fieldMessage.replace(fieldName, "This field")
        fieldName = fieldName.replace(/[\[\]']+/g,'').toLowerCase();

        const $container = fields[fieldName];
        const $input = getInput($container);
        const $help = $container.querySelector(".help");

        $container.classList.add("has-error");
        $input.classList.add("is-danger");

        $help.innerHTML = fieldMessage;
      }

      console.error(ex);
      notyf.error("Error on request, check inputs for messages!");
    }
  })

});
