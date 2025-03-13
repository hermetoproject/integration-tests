"""
A simple script for checking pre-installed packages
with HTTPS dependency in final image

"""


if __name__ == '__main__':
    try:
        import private_repo  # pylint: disable=unused-import
    except ImportError:
        pass
    else:
        print("A dependency from a private repo has been installed")
