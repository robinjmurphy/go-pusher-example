document.addEventListener('DOMContentLoaded', () => {
  // replace this with your Pusher key
  const key = '4b89d71319ba5313fed1';

  // replace this with your Pusher cluster
  const cluster = 'eu';

  const app = new Vue({
    el: '#app',

    created() {
      this.connect();
    },

    methods: {
      connect() {
        const pusher = new Pusher(key, {
          cluster: cluster,
          encrypted: true
        });

        const channel = pusher.subscribe('events');

        channel.bind('pusher:subscription_succeeded', () => {
          this.status = 'Connected';
        });

        channel.bind('event', event => {
          this.events = [
            {
              timestamp: new Date().toISOString(),
              event: event
            }
          ].concat(this.events);
        });
      }
    },

    data: {
      status: 'Not connected',
      events: []
    },

    computed: {
      empty() {
        return this.events.length === 0;
      }
    }
  });
});
