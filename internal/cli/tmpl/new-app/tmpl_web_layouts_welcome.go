package newapp

const tmplWebLayoutsWelcome string = `package layouts

import "time"

templ Welcome() {
	@Page("Welcome"){
    <div class="relative flex flex-col items-center justify-center min-h-screen px-4 overflow-hidden bg-white bg-opacity-95">
      <div class="absolute inset-0 z-0">
        <div class="border-t-8 border-blue-600 bg-white"></div>
        <div class="border-t-8 border-yellow-300 bg-white"></div>
        <div class="border-t-8 border-green-600 bg-white"></div>
      </div>

      <div class="absolute inset-0 z-0">
        <div
          class="absolute top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2 w-[560px] h-[560px] bg-[var(--secondary-glow)] blur-2xl opacity-50 rounded-full">
        </div>
        <div
          class="absolute top-1/2 left-1/2 -translate-x-1/3 -translate-y-1/2 w-[560px] h-[560px] bg-[var(--primary-glow)] blur-2xl opacity-30">
        </div>
      </div>

      <main class="relative z-10 flex flex-col items-center justify-center w-full max-w-5xl text-center">
        <div class="">
          <img alt="Jangada Framework Logo" class="h-64 w-auto" src="public/background.png" />
        </div>
        <div class="mx-8">
          <h1 class="mb-3 text-4xl font-bold text-gray-900">Jangada Framework</h1>
          <p class="text-lg font-medium text-gray-600 mb-10">
            A full-stack web framework in Go for building modern web applications, RESTful APIs, and gRPC with integrated frontend and backend.
          </p>
        </div>
      </main>
      
      <footer class="absolute bottom-4 text-center text-white-600 w-full">
        <p class="mt-3 text-gray-900 text-silver-500">
          <span class="mx-2">Jangada v0.1.0-beta</span> - 
          <span class="mx-2">Go v1.22.4</span> - 
          <span class="mx-2">{ time.Now().Format(time.RFC1123) }</span>
        </p>
      </footer>

    </div>
  }
}
`
