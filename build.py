import os
import subprocess

FUNCTIONS = [
    'hello',
    'world'
]

def main():
    os.environ['GOOS'] = 'linux'
    os.environ['GOARCH'] = 'amd64'
    os.environ['CGO_ENABLED'] = '0'
    for func in FUNCTIONS:
        print(f'building {func}')
        result = subprocess.run(["go", "build", "-v" ,"-o", f"bin/{func}", f"{func}/main.go"], capture_output=True)
        print(result.stdout.decode())
        print(result.stderr.decode())


if __name__ == "__main__":
    main()