import os


def main():
    file_path = os.path.abspath("Go/main.go")
    with open(file_path, "r", encoding="utf-8") as f:
        code = f.read()

    
    
if __name__ == "__main__":
    main()