#include <iostream>
using namespace std;



int main(){
	int *p ; // wild pointer
	cout << "address of wild pointer is " << p << endl;
	cin >> *p;
	cout <<"enterred value is :" << *p << endl;

	int* p2 = NULL; // points to nothing
	cout << "value of p2 is (p2 is pointing to) "<< p2 << endl;
	cout <<  "value of *p2 is " << *p2 << endl;
	return 0;
}
