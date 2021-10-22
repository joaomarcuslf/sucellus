document.addEventListener('DOMContentLoaded', async () => {
  const ID = window.location.href.substring(window.location.href.lastIndexOf('/') + 1);

  const $backBtn = document.querySelector("#backBtn");
  $backBtn.addEventListener('click', async () => {
    window.history.back();
  });

  const $stopBtn = document.querySelector("#stopBtn");
  $stopBtn && $stopBtn.addEventListener('click', async () => {
    try {
      await postJson(`/api/services/${ID}/stop`)

      notyf.success("Stopped service!");

      window.location.reload(false);
    } catch(ex) {
      console.error(ex);
      notyf.error("Error stopping service!");
    }
  });

  const $startBtn = document.querySelector("#startBtn");
  $startBtn && $startBtn.addEventListener('click', async () => {
    try {
      await postJson(`/api/services/${ID}/start`)

      notyf.success("Stopped service!");

      window.location.reload(false);
    } catch(ex) {
      console.error(ex);
      notyf.error("Error stopping service!");
    }
  });

  const $updateBtn = document.querySelector("#updateBtn");
  $updateBtn.addEventListener('click', async () => {
    window.location.href = `/services/${ID}/update`;
  });

  const $deleteBtn = document.querySelector("#deleteBtn");
  $deleteBtn.addEventListener('click', async () => {
    try {
      await deleteJson(`/api/services/${ID}`)

      notyf.success("Deleted service!");

      window.location.href = "/services";
    } catch(ex) {
      console.error(ex);
      notyf.error("Error deleting service!");
    }
  });
});
