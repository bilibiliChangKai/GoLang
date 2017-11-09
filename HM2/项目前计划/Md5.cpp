#include <iostream>
#include <string>
#include <cmath>
#include <sstream>
#include <string>
#include <iomanip>
#include <stdio.h>
using namespace std;
#define MAX_UNSIGNED_INT 4294967295

// ʮ�������ַ�������
char hexCh[16] = {'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f'};
// A-D�ĸ�32λ�Ĵ���
unsigned A = 0x67452301;
unsigned B = 0xEFCDAB89;
unsigned C = 0x98BADCFE;
unsigned D = 0x10325476;
// s����
unsigned s[64] = {7, 12, 17, 22, 7, 12, 17, 22, 7, 12, 17, 22, 7, 12, 17, 22,
                  5, 9,  14, 20, 5, 9,  14, 20, 5, 9,  14, 20, 5, 9,  14, 20,
                  4, 11, 16, 23, 4, 11, 16, 23, 4, 11, 16, 23, 4, 11, 16, 23,
                  6, 10, 15, 21, 6, 10, 15, 21, 6, 10, 15, 21, 6, 10, 15, 21};
// ��¼�ַ�����
unsigned lengthB = 0, lengthS = 0;

// �ĸ�����
unsigned F(unsigned X, unsigned Y,unsigned Z) {
    return (X & Y) | ((~X) & Z);
}
unsigned G(unsigned X, unsigned Y, unsigned Z) {
    return (X & Z) | (Y & (~Z));
}
unsigned H(unsigned X, unsigned Y, unsigned Z) {
    return X ^ Y ^ Z;
}
unsigned I(unsigned X, unsigned Y, unsigned Z) {
    return Y ^ (X | (~Z));
}

// T��������1��ʼ
unsigned T(unsigned i) {
    return (MAX_UNSIGNED_INT + 1) * abs(sin(i));
}

// �ĸ�X����
int X1(int i) {
    return i;
}
int X2(int i) {
    return (1 + 5 * i) % 16;
}
int X3(int i) {
    return (5 + 3 * i) % 16;
}
int X4(int i) {
    return 7 * i % 16;
}

// ��sѭ������numλ
unsigned CLS(unsigned s, int num) {
    return (s << num) | (s >> (32 - num));
}

void addMessageLength(int i) {
    // ������������λ
    if (lengthS + i < lengthS)  lengthB++;
    lengthS += i;
}

// MD5hash����
// input��CVi-1(4*32)��Y(16*32)
// output��CVi_1 -> CVi(4*32)
// Position: 0 1 2 3 4 .. 16
void H_MD5(unsigned *CVi_1, unsigned *Y) {
    A = CVi_1[0];
    B = CVi_1[1];
    C = CVi_1[2];
    D = CVi_1[3];
    unsigned temp;

    // Round 1
    for (int i = 0; i < 16; i++) {
        temp = B + CLS(A + F(B, C, D) + Y[X1(i)] + T(i + 1), s[i]);
        A = D;
        D = C;
        C = B;
        B = temp;
    }

    // Round 2
    for (int i = 0; i < 16; i++) {
        temp = B + CLS((A + G(B, C, D) + Y[X2(i)] + T(i + 17)), s[i + 16]);
        A = D;
        D = C;
        C = B;
        B = temp;
    }

    // Round 3
    for (int i = 0; i < 16; i++) {
        temp = B + CLS((A + H(B, C, D) + Y[X3(i)] + T(i + 33)), s[i + 32]);
        A = D;
        D = C;
        C = B;
        B = temp;
    }

    // Round 4
    for (int i = 0; i < 16; i++) {
        temp = B + CLS((A + I(B, C, D) + Y[X4(i)] + T(i + 49)), s[i + 48]);
        A = D;
        D = C;
        C = B;
        B = temp;
    }

    CVi_1[0] += A;
    CVi_1[1] += B;
    CVi_1[2] += C;
    CVi_1[3] += D;
}

// �����ַ�����ֱ��64λ����EOFΪֹ
string inputBuffer() {
    char input;
    string str;
    for (int i = 0; i < 64; i++) {
        if (fscanf(stdin, "%c", &input) == -1)
            break;
        addMessageLength(8);
        str += input;
    }
    return str;
}

void addPadding(string &input) {
    // ����Ҫ��1
    input += (char)0x80;
    int len = input.length();
    // ��������1�󣬳��ȴ���56������������һ��
    for (int i = len; i < 56 + 64 * (len > 56); i++)
        input += (char)0x0;
}

// ����512��bit(64��char)��16��unsigned int
unsigned *processStringToInts(string message) {
    unsigned *returnInts = new unsigned[16];
    for (int i = 0; i < 16; i++)  returnInts[i] = 0;
    // ����ǰ14λ
    for (int i = 0; i < 14; i++)
        for (int j = 0; j < 4; j++)
            returnInts[i] = (returnInts[i] << 8) + ((unsigned)(message[i * 4 + 3 - j]) & 0xFF);
    // ����������λ
    if (message.length() == 64)
        for (int i = 14; i < 16; i++)
            for (int j = 0; j < 4; j++)
                returnInts[i] = (returnInts[i] << 8) + ((unsigned)(message[i * 4 + 3 - j]) & 0xFF);
    else {
        returnInts[14] = lengthS;
        returnInts[15] = lengthB;
    }
    return returnInts;
}

// ��8λintת��16�����ַ���
string decToHex(unsigned input)
{
    string temp = "";     //����ת�����ַ�
    temp += hexCh[(input / (1 << 4)) - 0];
    temp += hexCh[(input % (1 << 4)) - 0];
    return temp;
}

// ��CV[4]����С��ת�����ַ���
string processIntsToString(unsigned *CV) {
    string returnStr = "";
    unsigned temp[4];
    for (int i = 0; i < 4; i++) {
        temp[i] = CV[i];
        for (int j = 0; j < 4; j++) {
            returnStr += decToHex(temp[i] % (1 << 8));
            temp[i] = temp[i] >> 8;
        }
    }
    return returnStr;
}

int main() {
    string input, output;
    bool endFlag = false;
    unsigned *Y, *CV;

    // CV��ʼ��
    CV = new unsigned[4];
    CV[0] = A;
    CV[1] = B;
    CV[2] = C;
    CV[3] = D;

    while (!endFlag) {
        input = inputBuffer();
        // �����Ѿ�����
        if (input.length() < 64) {
            endFlag = true;
            addPadding(input);
            // �����������������飬�Ƚ���һ��Hash
            if (input.length() != 56) {
                Y = processStringToInts(input.substr(0, 64));
                H_MD5(CV, Y);
                input.erase(0, 64);
            }
        }
        // ����
        Y = processStringToInts(input);

        H_MD5(CV, Y);
        delete[] Y;
    }

    output = processIntsToString(CV);
    fprintf(stdout, "%s", output.c_str());
}
