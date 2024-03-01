"""
题目描述

公司用一个字符串来表示员工的出勤信息

absent: 缺勤
late: 迟到
leaveearly: 早退
present: 正常上班.
现需根据员工出勤信息，判断本次是否能获得出勤奖，能获得出勤奖的条件如下

缺勤不超过一次
没有连续的迟到/早退,
任意连续 7 次考勤，缺勤/迟到/早退不超过 3 次
输入描述

用户的考勤数据字符串

记录条数 >= 1;

输入字符串长度< 10000:

不存在非法输入;
"""


def calc(infos: list[str]) -> bool:
    absent_count = 0
    is_pre_late = False
    err_count = 0

    count = 0
    for info in infos:
        count += 1
        if info == "absent":
            if absent_count > 1:
                return False
            absent_count += 1
            err_count += 1
            is_pre_late = False
        elif info in {"late", "leaveearly"}:
            if is_pre_late:
                return False
            err_count += 1
            is_pre_late = True
        elif info == "present":
            is_pre_late = False

        if count == 7:
            if err_count >= 3:
                return False
            err_count = 0

    return True


if __name__ == "__main__":
    print(calc(["present"]))
    print(
        calc(
            [
                "present",
                "absent",
                "present",
                "present",
                "leaveearly",
                "present",
                "absent",
            ]
        )
    )
