<script>
  import { onMount } from "svelte";
  import logo from "./assets/images/logo-universal.png";
  import { EventsOn } from "../wailsjs/runtime/runtime.js";
  import { StartTimer, StopTimer, GetEvents } from "../wailsjs/go/main/App.js";

  let isRunning = false;
  let remainingSeconds = 0;
  let timerDisplay = "0:00";
  let events = [];

  async function loadEvents() {
    events = await GetEvents();
  }

  onMount(() => {
    loadEvents();

    EventsOn("timer:tick", (seconds) => {
      remainingSeconds = seconds;
      const mins = Math.floor(seconds / 60);
      const secs = seconds % 60;
      timerDisplay = `${mins}:${secs.toString().padStart(2, "0")}`;
    });

    EventsOn("timer:complete", () => {
      isRunning = false;
      remainingSeconds = 0;
      timerDisplay = "0:00";
      loadEvents();
      console.log("Timer complete!");
    });
  });

  function toggleTimer() {
    if (isRunning) {
      StopTimer();
    } else {
      StartTimer();
    }
    isRunning = !isRunning;
  }
</script>

<main>
  <img alt="Wails logo" id="logo" src={logo} />
  <div class="timer-display">{timerDisplay}</div>
  <div class="status">{isRunning ? "Running..." : "Stopped"}</div>
  <div class="input-box" id="input">
    <button class="btn" on:click={toggleTimer}>
      {isRunning ? "Stop" : "Start"}
    </button>
  </div>

  <div class="events-section">
    <h2>Pomodoro Events</h2>
    {#if events.length === 0}
      <p>No events yet. Start a timer!</p>
    {:else}
      <ul class="events-list">
        {#each events as event}
          <li class="event-item">
            <div class="event-name">{event.Name}</div>
            <div class="event-time">
              Started: {new Date(event.StartTime).toLocaleTimeString()}
              {#if event.EndTime && event.EndTime !== "0001-01-01T00:00:00Z"}
                - Ended: {new Date(event.EndTime).toLocaleTimeString()}
              {/if}
            </div>
          </li>
        {/each}
      </ul>
    {/if}
  </div>
</main>

<style>
  #logo {
    display: block;
    width: 50%;
    height: 50%;
    margin: auto;
    padding: 10% 0 0;
    background-position: center;
    background-repeat: no-repeat;
    background-size: 100% 100%;
    background-origin: content-box;
  }

  .timer-display {
    font-size: 3rem;
    font-weight: bold;
    margin: 2rem auto;
    text-align: center;
  }

  .status {
    height: 20px;
    line-height: 20px;
    margin: 1rem auto;
    text-align: center;
  }

  .input-box .btn {
    width: 60px;
    height: 30px;
    line-height: 30px;
    border-radius: 3px;
    border: none;
    margin: 0 0 0 20px;
    padding: 0 8px;
    cursor: pointer;
  }

  .input-box .btn:hover {
    background-image: linear-gradient(to top, #cfd9df 0%, #e2ebf0 100%);
    color: #333333;
  }

  .input-box .input {
    border: none;
    border-radius: 3px;
    outline: none;
    height: 30px;
    line-height: 30px;
    padding: 0 10px;
    background-color: rgba(240, 240, 240, 1);
    -webkit-font-smoothing: antialiased;
  }

  .input-box .input:hover {
    border: none;
    background-color: rgba(255, 255, 255, 1);
  }

  .input-box .input:focus {
    border: none;
    background-color: rgba(255, 255, 255, 1);
  }

  .events-section {
    margin: 2rem auto;
    max-width: 600px;
    padding: 0 1rem;
  }

  .events-section h2 {
    text-align: center;
    margin-bottom: 1rem;
  }

  .events-list {
    list-style: none;
    padding: 0;
  }

  .event-item {
    background-color: rgba(240, 240, 240, 0.5);
    border-radius: 5px;
    padding: 1rem;
    margin-bottom: 0.5rem;
  }

  .event-name {
    font-weight: bold;
    margin-bottom: 0.5rem;
  }

  .event-time {
    font-size: 0.9rem;
    color: #666;
  }
</style>
