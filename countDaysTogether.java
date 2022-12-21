class Solution {
    public static int[] numOfDaysPerMonth = new int[] {31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31};

    public int countDaysTogether(String arriveAlice, String leaveAlice, String arriveBob, String leaveBob) {

        int aliceArriveMonth = getMonth(arriveAlice);
        int aliceArriveDay = getDay(arriveAlice);
        
        int aliceLeaveMonth = getMonth(leaveAlice);
        int aliceLeaveDay = getDay(leaveAlice);

        int bobArriveMonth = getMonth(arriveBob);
        int bobArriveDay = getDay(arriveBob);
        
        int bobLeaveMonth = getMonth(leaveBob);
        int bobLeaveDay = getDay(leaveBob);

        int bobArrivalDay = numOfDayInYear(bobArriveDay, bobArriveMonth);
        int bobLeavingDay = numOfDayInYear(bobLeaveDay, bobLeaveMonth);

        int aliceArrivalDay = numOfDayInYear(aliceArriveDay, aliceArriveMonth);
        int aliceLeavingDay = numOfDayInYear(aliceLeaveDay, aliceLeaveMonth);

        if (aliceLeavingDay < bobArrivalDay || bobLeavingDay < aliceArrivalDay) {
            return 0;
        }

        if (bobArrivalDay > aliceArrivalDay ) {
            if (aliceLeavingDay < bobLeavingDay) {
                return aliceLeavingDay - bobArrivalDay + 1;
            } else if (bobLeavingDay <= aliceLeavingDay) {
                return bobLeavingDay - bobArrivalDay + 1;
            }
        }

        if (bobArrivalDay <= aliceArrivalDay) {
            if (bobLeavingDay < aliceLeavingDay) {
                return bobLeavingDay - aliceArrivalDay + 1;
            } else if (bobLeavingDay >= aliceLeavingDay) {
                return aliceLeavingDay - aliceArrivalDay + 1;
            }
        }
        
        return 0;
    }

    private int numOfDayInYear(int day, int month) {
        int result = day;

        for (int i = 0; i < month - 1; i++) {
            result += numOfDaysPerMonth[i];
        }

        return result;
    }

    private int getDay(String date) {
        return Integer.parseInt(date.substring(3,5));
    }

    private int getMonth(String date) {
        return Integer.parseInt(date.substring(0,2));
    }
}
