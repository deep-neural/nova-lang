<h1 align="center">
  <a href="https://github.com/deep-neural/nova-lang"><img src="./.github/logo.jpg" alt="Nova Lang" height="150px"></a>
  <br>
  Nova Lang
  <br>
</h1>
<h4 align="center">A modern AI-first programming language with memory safety and concurrency at its core</h4>
<p align="center">
    <a href="https://github.com/deep-neural/nova-lang"><img src="https://img.shields.io/badge/NovaLang-AI-5865F2.svg?longCache=true" alt="Nova Lang" /></a>
  <a href="https://github.com/deep-neural/nova-lang"><img src="https://img.shields.io/static/v1?label=Version&message=1.0.0&color=brightgreen" /></a>
  <a href="https://github.com/deep-neural/nova-lang"><img src="https://img.shields.io/static/v1?label=Safety&message=Memory-Safe&color=brightgreen" /></a>
  <a href="https://github.com/deep-neural/nova-lang"><img src="https://img.shields.io/static/v1?label=Paradigm&message=Multi-Paradigm&color=brightgreen" /></a>
  <a href="https://github.com/deep-neural/nova-lang"><img src="https://img.shields.io/static/v1?label=Release&message=Stable&color=brightgreen" /></a>
  <br>
    <a href="https://github.com/deep-neural/nova-lang"><img src="https://img.shields.io/static/v1?label=Build&message=Documentation&color=brightgreen" /></a>
    <a href="LICENSE"><img src="https://img.shields.io/badge/License-MIT-5865F2.svg" alt="License: MIT" /></a>
</p>
<br>

### New Release

Nova Lang v1.0.0 has been released! See the [release notes](https://github.com/novalang/nova/wiki/Release-NovaLang@v1.0.0) to learn about new features, enhancements, and breaking changes.

If you aren't ready to upgrade yet, check the [tags](https://github.com/novalang/nova/tags) for previous stable releases.

We appreciate your feedback! Feel free to open GitHub issues or submit changes to stay updated in development and connect with the maintainers.

-----

### Usage

Nova is distributed as a compiler toolchain with a comprehensive standard library. To integrate it into your project, ensure you have the Nova compiler (`nova`) and the necessary build tools (e.g., Nova Build System, NBS). Clone the repository and use the NBS to manage your project dependencies and build process.

## Nova Compiler
Nova translates source to LLVM IR, then compiles to machine code. 
```bash
$ nova main.nv -o main

$ ./main
```

## Quick Start

```nova
import "foundation"

func main() {
    print("Hello, Nova!");
}
```

## AI-First Development
Nova Lang revolutionizes development with first-class AI integration:

```nova
// Import AI models directly
import "gemma7b" from "generative/ai/models/text"
import "vision" from "generative/ai/models/image"

func main() {
    // Simple prompt-based generation
    response = gemma.generate("Explain quantum computing in simple terms");
    print(response);
    
    // Computer vision integration
    image = vision.loadImage("photo.jpg");
    objects = vision.detectObjects(image);
    print("Detected", objects.length, "objects in the image");
    
    // Process image with specific parameters
    analyzed = vision.analyze(image, {
        detail_level: "high",
        detect: ["faces", "text", "landmarks"],
        confidence_threshold: 0.85
    });
    
    for item in analyzed.detections {
        print(item.type, "at position", item.bounds, "confidence:", item.confidence);
    }
}
```

## Virtual Machine Blocks
Execute code in isolated environments with precise dependency control:

```nova
import "llama" from "models/llm"

func main() {
    // Define a virtual environment with specific dependencies
    virtual func() {
        // Run LLM inference in isolated environment with GPU acceleration
        model = llama.load("llama-3-70b");
        response = model.generate("Design a sorting algorithm that works well for nearly-sorted data");
        
        // Stream results back to main process
        stream, send response
    }({
        gpu: true,
        memory: "16GB",
        cuda_version: "12.0"
    })
    
    // Cloud provider integration
    config = {
        provider: "google-cloud",
        service_account: "service-account.json",
        machine_type: "a2-highgpu-1g",
        gpu_type: "nvidia-a100"
    }
    
    virtual func() {
        // Run large-scale training job
        trainer = ml.createTrainer("my_model", {
            epochs: 100,
            batch_size: 64,
            learning_rate: 0.001
        });
        
        result = trainer.train();
        stream, send result
    }(config)
}
```

## Native C/C++ Interoperability
Directly use low-level libraries with zero overhead:

```nova
import "stdio" from "extern"

struct Point { x: i32, y: i32 }

func main() {
    // Direct C function calls
    var name = "Nova";
    stdio.printf("Hello from %s using C printf!\n", name.to_c_string());
    
    // File operations using C API
    var file = stdio.fopen("data.txt".to_c_string(), "w".to_c_string());
    if file == stdio.NULL {
        stdio.perror("Failed to open file");
        return;
    }
    
    let p = Point { x: 10, y: 20 };
    stdio.fprintf(file, "Point data: (%d, %d)\n", p.x, p.y);
    stdio.fclose(file);
}
```

## UI and State Management
Build reactive UIs with native and web components:

```nova
import { webview, component } from "ui"

func UserProfile() -> component {
    // Local component state
    var [name, setName] = ui.State("Guest");
    
    // Remote state (syncs with backend automatically)
    var [userData, setUserData] = ui.RemoteState("/api/user");
    
    ui.Effect(() => {
        if userData.loaded {
            setName(userData.fullName);
        }
    }, [userData]);
    
    return (
        <div class="profile">
            <h1>Welcome, {name}</h1>
            <button onClick={() => setName("New Name")}>
                Update Name
            </button>
            
            <div class="stats">
                <p>Member since: {userData.joinDate}</p>
                <p>Last login: {userData.lastLogin}</p>
            </div>
        </div>
    );
}

func main() {
    webview.render(UserProfile());
}
```

## Package Manager

```bash
// Virtual environments
$ npkg install vision-model --virtual

// Local environment
$ npkg install llama-model

// Remote repositories
$ npkg install https://github.com/user/custom-model.git
```

## Import Flexibility
Multiple import styles for different use cases:

```nova
// Standard imports
import { generate, embedText } from "llm"

// Remote imports with version control
import "stable-diffusion" from "huggingface.co/models" (
    version = "3.0",
    cache = true
)

// GitHub imports with authentication
import "private-model" from "github.com/organization" (
    branch = "main",
    protocol = "https",
    auth = "bearer",
    token = env.GITHUB_TOKEN,
    verify = "sha256:abc123..."
)

// Local file imports
import { MyComponent } from "./components"
```

### Features

#### Modern Type System
* Strong static typing with powerful type inference for concise code
* Algebraic data types with pattern matching for expressive data handling
* Generics with specialization for type-safe generic programming
* Traits and interfaces for flexible abstraction and composition
* Immutability by default for safer concurrent programming

#### AI Integration
* First-class AI model imports and integration
* Efficient inference runtime with hardware acceleration
* Type-safe model inputs and outputs
* Runtime model optimization and quantization
* Unified interface for multiple model architectures

#### Memory Safety & Management
* Region-based memory management with compile-time lifetime analysis
* Zero-cost abstractions for memory-safe systems programming
* Optional garbage collection for domains where appropriate
* Smart pointers with ownership semantics to prevent memory leaks
* Compile-time checking of memory access patterns

#### Concurrency & Parallelism
* Built-in async/await syntax for intuitive asynchronous programming
* Advanced task scheduler optimized for many-core architectures
* Actor model for safe concurrent state management
* Software transactional memory (STM) for composable concurrency
* Lock-free data structures in the standard library

#### Performance Optimizations
* LLVM-based optimizing compiler with whole-program optimization
* Profile-guided optimization (PGO) for real-world performance gains
* Fine-grained control over memory layout and alignment
* SIMD intrinsics and automatic vectorization
* Compile-time function execution and metaprogramming

#### Isolated Execution Environments
* Virtual blocks for dependency isolation
* Cloud provider integration for compute-intensive tasks
* Resource management for GPU and specialized hardware
* Transparent IPC between virtual environments
* Configurable security and privilege boundaries

#### Development Experience
* Comprehensive toolchain with integrated package manager
* Incremental compilation for fast development cycles
* Built-in formatting, linting, and documentation generation
* Property-based testing and fuzzing integrations
* Language server protocol support for IDE integration

#### Nova Build System (NBS)
* Declarative build configurations with dependency resolution
* Reproducible builds with hermetic build environments
* Distributed build cache for fast CI/CD integration
* Cross-compilation support for multiple target architectures
* Integrated benchmark and performance analysis tools

### Contributing

Check out the [contributing guide](https://github.com/novalang/nova/wiki/Contributing) to join the team of dedicated contributors making this project possible.

### License

MIT License - see [LICENSE](LICENSE) for full text
