function by_id(id_string) {
  return document.getElementById(id_string);
}

function draw_user_button() {
  by_id('clerk-app').innerHTML = `<div id="user-button"></div>`;
  Clerk.mountUserButton(by_id('user-button'))
} // func


function draw_sign_in_button() {
  by_id('clerk-app').innerHTML = `<div id="sign-in"></div>`;
  Clerk.mountSignIn(by_id('sign-in'))
}

      window.addEventListener('load', async function () {
        await Clerk.load()

        console.log('ClerkJS is loaded via CDN. 1')

        if (Clerk.isSignedIn) {
          console.log('drawing user-button')
          draw_user_button();
        } else {
          console.log('drawing sign-in')
          draw_sign_in_button();
          console.log('done sign-in')
        }

      }); // Window Event
