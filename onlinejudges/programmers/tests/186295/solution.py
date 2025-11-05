# https://programmers.co.kr/app/with_setting/tests/186295/challenges

import random
import hashlib


def hash(item):
    return hashlib.sha1(str(item).encode("utf-8")).hexdigest()


def fingerprint(item):
    return int(hash(item)[:8], 16)


class Meet:
    capacity = 1 << 8
    buckets = [[] for _ in range(1 << 8)]
    max_kicks = 100
    bucket_size = 4
    size = 0

    @classmethod
    def where(cls, number):
        fp = fingerprint(number)
        h1 = int(hash(number), 16) % cls.capacity
        h2 = (h1 ^ fp) % cls.capacity
        return fp, h1, h2

    @classmethod
    def meet(cls, number):
        fp, h1, h2 = cls.where(number)

        if len(cls.buckets[h1]) < cls.bucket_size:
            cls.buckets[h1].append(fp)
            cls.size += 1
            return True

        if len(cls.buckets[h2]) < cls.bucket_size:
            cls.buckets[h2].append(fp)
            cls.size += 1
            return True

        i = random.choice([h1, h2])
        for _ in range(cls.max_kicks):
            kick_index = random.randint(0, cls.bucket_size - 1)
            kick_fp = cls.buckets[i].pop(kick_index)

            cls.buckets[i].append(fp)
            fp = kick_fp

            i = (i ^ fp) % cls.capacity
            if len(cls.buckets[i]) < cls.bucket_size:
                cls.buckets[i].append(fp)
                cls.size += 1
                return True

            return False

    @classmethod
    def check(cls, number):
        fp, h1, h2 = cls.where(number)

        if fp in cls.buckets[h1]:
            return True
        if fp in cls.buckets[h2]:
            return True
        return False

    @classmethod
    def clear(cls):
        cls.buckets = [[] for _ in range(cls.capacity)]
        cls.size = 0
