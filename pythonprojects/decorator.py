import logging
from contextlib import contextmanager

@contextmanager
def debug_logging(level):
    logger = logging.getLogger()
